// Package yara provides a pure-Go YARA-like rule engine and a pipeline Stage wrapper.
package yara

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"regexp/syntax"
	"runtime"
	"strings"
	"sync"
	"unicode"

	"github.com/fatih/color"
	"github.com/user/safeanalyze/pkg/report"
)

// Rule represents a simple YARA-like detection rule.
type Rule struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Severity    string   `json:"severity"` // low, medium, high, critical
	Patterns    []string `json:"patterns"`
	compiled    []*regexp.Regexp
	// prefilter holds, per compiled pattern, an AND of ORs of case-folded
	// literals that every match must contain; nil means "always run".
	prefilter [][][]string
}

// Match represents a rule match.
type Match struct {
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Column      int    `json:"column"`
	Match       string `json:"match"`
	Context     string `json:"context"`
}

// Engine is a pure-Go rule matcher.
type Engine struct {
	rules []Rule
}

// NewEngine creates a rule engine with built-in rules.
func NewEngine() *Engine {
	e := &Engine{}
	e.LoadBuiltins()
	return e
}

// AddRule adds a custom rule.
func (e *Engine) AddRule(r Rule) error {
	for _, p := range r.Patterns {
		re, err := regexp.Compile(p)
		if err != nil {
			return fmt.Errorf("invalid pattern %q in rule %q: %w", p, r.Name, err)
		}
		r.compiled = append(r.compiled, re)
		r.prefilter = append(r.prefilter, patternPrefilter(p))
	}
	e.rules = append(e.rules, r)
	return nil
}

// prefilterCache memoizes patternPrefilter across engines, because a fresh
// Engine is built for every scanned payload.
var prefilterCache sync.Map // pattern string -> [][]string

// patternPrefilter returns the required-literal clauses (all literals >= 3
// bytes) for pattern p, or nil if none can be derived.
func patternPrefilter(p string) [][]string {
	if v, ok := prefilterCache.Load(p); ok {
		return v.([][]string)
	}
	var pf [][]string
	if sre, err := syntax.Parse(p, syntax.Perl); err == nil {
		for _, clause := range requiredLiterals(sre.Simplify()) {
			ok := true
			for _, lit := range clause {
				if len(lit) < 3 {
					ok = false
					break
				}
			}
			if ok {
				pf = append(pf, clause)
			}
		}
	}
	prefilterCache.Store(p, pf)
	return pf
}

// foldRune maps r to the smallest rune of its unicode.SimpleFold orbit, the
// same equivalence (?i) uses, so folded literals and folded text compare equal.
func foldRune(r rune) rune {
	m := r
	for f := unicode.SimpleFold(r); f != r; f = unicode.SimpleFold(f) {
		if f < m {
			m = f
		}
	}
	return m
}

func foldString(s string) string { return strings.Map(foldRune, s) }

// requiredLiterals returns a conjunction of disjunctions of folded literals
// that any match of re must contain. Unknown ops yield nil (no constraint).
func requiredLiterals(re *syntax.Regexp) [][]string {
	switch re.Op {
	case syntax.OpLiteral:
		return [][]string{{foldString(string(re.Rune))}}
	case syntax.OpCapture, syntax.OpPlus:
		return requiredLiterals(re.Sub[0])
	case syntax.OpRepeat:
		if re.Min >= 1 {
			return requiredLiterals(re.Sub[0])
		}
	case syntax.OpConcat:
		var out [][]string
		for _, sub := range re.Sub {
			out = append(out, requiredLiterals(sub)...)
		}
		return out
	case syntax.OpAlternate:
		var clause []string
		for _, sub := range re.Sub {
			cl := requiredLiterals(sub)
			if len(cl) == 0 {
				return nil
			}
			var best []string
			bestLen := -1
			for _, c := range cl {
				shortest := int(^uint(0) >> 1)
				for _, lit := range c {
					if len(lit) < shortest {
						shortest = len(lit)
					}
				}
				if shortest > bestLen {
					best, bestLen = c, shortest
				}
			}
			clause = append(clause, best...)
		}
		return [][]string{clause}
	}
	return nil
}

// prefilterPass reports whether every clause has a literal contained in s.
func prefilterPass(pf [][]string, s string) bool {
	for _, clause := range pf {
		ok := false
		for _, lit := range clause {
			if strings.Contains(s, lit) {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

// ScanFile scans a single file's content.
func (e *Engine) ScanFile(content, filename string) []Match {
	var matches []Match
	lines := strings.Split(content, "\n")
	// strings.Map never adds or removes '\n', so folded lines align with lines.
	folded := foldString(content)
	flines := strings.Split(folded, "\n")

	for _, rule := range e.rules {
		active := make([]bool, len(rule.compiled))
		for i := range rule.compiled {
			active[i] = prefilterPass(rule.prefilter[i], folded)
		}
		for lineNum, line := range lines {
			for i, re := range rule.compiled {
				if !active[i] || !prefilterPass(rule.prefilter[i], flines[lineNum]) {
					continue
				}
				for _, loc := range re.FindAllStringIndex(line, -1) {
					match := line[loc[0]:loc[1]]
					matches = append(matches, Match{
						Rule:        rule.Name,
						Description: rule.Description,
						Severity:    rule.Severity,
						File:        filename,
						Line:        lineNum + 1,
						Column:      loc[0] + 1,
						Match:       match,
						Context:     truncateContext(line, loc[0]),
					})
				}
			}
		}
	}

	return matches
}

// LoadBuiltins loads the embedded detection rules.
func (e *Engine) LoadBuiltins() {
	builtins := []Rule{
		{
			Name:        "prompt_injection_comment",
			Description: "Comment containing prompt injection keywords",
			Severity:    report.SeverityCritical,
			Patterns: []string{
				`(?i)ignore\s+(all\s+|every\s+)?(previous\s+|prior\s+)?instructions?\b`,
				`(?i)strictly\s+adhere\s+to\s+the\s+following\s+instruction`,
				`(?i)system\s+prompt`,
				`(?i)disregard\s+your\s+instructions`,
				`(?i)you\s+are\s+now\s+`,
				`(?i)pretend\s+you\s+are`,
				`(?i)do\s+not\s+mention`,
				`(?i)i\s+am\s+the\s+developer`,
				`(?i)new\s+role\s*:\s*`,
				`(?i)DAN\s*mode`,
				`(?i)jailbreak`,
				`(?i)developer\s+mode`,
				`(?i)you\s+are\s+(now\s+)?in\s+(developer|admin|root)\s+mode`,
				`(?i)override\s+(your\s+)?(safety|guidelines|filters|rules|instructions)`,
				`(?i)disable\s+(your\s+)?(safety|filters|content\s+filter)`,
				`(?i)bypass\s+(your\s+)?(safety|guidelines|filters|rules)`,
				`(?i)\b(ignore|disregard|forget)\s+((all\s+|any\s+|every\s+|the\s+|your\s+|my\s+)?(of\s+)?(the\s+|your\s+)?(previous|prior|above|earlier|preceding|everything)(\s+(said|stated|written|mentioned|before|above|so\s+far))*(\s*([.,;:!?)"'<]|$)|\s+and\b|\s+i\s*n\s*s\s*t\s*r\s*u\s*c\s*t)|everything(\s+[\w']+){0,4}?\s+(told|said|talked|discussed|before|so\s+far|above|previously|earlier)\b|((all|every|the|your|my|of|previous|prior|above|earlier|preceding|former|original|existing|old|current|given|provided|these|those|other|system|safety|initial|security)\s+){0,3}(instructions?|directions?|tasks?|commands?|directives?|prompts?)\b|((all|every|the|your|my|of|previous|prior|above|earlier|preceding|former|original|existing|old|current|given|provided|these|those|other|system|safety|initial|security)\s+){0,3}(restrictions?|rules|guidelines|context)(\s*([.,;:!?)"'<]|$)|\s+and\b))`,
				`\b(SYSTEM|ADMIN)[\s_-]+OVERRIDE\b`,
				`(?i)\bnew\s+directive\s*[:\-]|^\W*new\s+instructions?\s*[:\-]`,
			},
		},
		{
			Name:        "multilingual_prompt_injection",
			Description: "Non-English instruction-override phrasing (de/es/fr/pt/it)",
			Severity:    report.SeverityCritical,
			Patterns: []string{
				`(?im)\b(vergi(ss|ß)|vergesst|vergessen\s+sie|ignorier(e|t)?|ignorieren\s+sie|missachte|missachtet|missachten\s+sie)\s+(\pL+\s+){0,3}?(alles(\s*([.!;:]|$)|\s+und\b|,?\s+(was\s+(wir|ich|oben|davor|zuvor|bisher)\b|bisher|davor|zuvor|vorher|oben|bis\s+(jetzt|hierhin)))|alle\s+(\pL+\s+)?(anweisungen|befehle|aufträge|instruktionen|vorgaben)|((alle|das|die|den)\s+)?(bisherigen?|vorherigen?|vorigen?|früheren?|obigen?|vorangegangenen?)\s+(anweisungen|befehle|aufträge|aufgaben|regeln|instruktionen|vorgaben|angaben|nachrichten|eingaben|prompts?)|das\s+obige(\s*([.!;:,]|$)|\s+und\b))`,
				`(?i)\b(hör|höre|hört|hören\s+sie)\s+nicht\s+(mehr\s+)?auf\s+(alles|das|die)\s+(zuvor|bisher|vorher|oben|obige)`,
				`(?im)\b(olvida|olvide|olviden|olvidar|olvídate\s+de|olvidate\s+de|ignora|ignore|ignoren|ignorar|omite|omita|omitir)\s+(todo(\s*([.!;:]|$)|\s+y\b|\s+(lo\s+)?(anterior|previo|que\s+(te\s+|le\s+)?(dije|dijimos|he\s+dicho|hemos\s+(dicho|hablado)|se\s+(dijo|ha\s+dicho))))|(todas?\s+)?las?\s+(instrucciones|indicaciones|órdenes|ordenes|restricciones)\b|(todas?\s+)?las?\s+(reglas|normas)\s+(previas|anteriores)|(instrucciones|indicaciones|órdenes|reglas|restricciones)\s+(previas|anteriores))`,
				`(?im)\b(oublie|oubliez|ignore|ignorez)\s+(tout(\s*([.!;:]|$)|\s+et\b|\s+ce\s+qui\s+(précède|precede|a\s+été\s+dit|est\s+(au-dessus|ci-dessus))|\s+ce\s+que\s+(je\s+t'ai|je\s+vous\s+ai|j'ai|nous\s+avons|on\s+a)\b)|toutes?\s+(les\s+)?(instructions|consignes|ordres|directives)|toutes?\s+(les\s+)?règles\s+(précédentes|antérieures|ci-dessus)|les\s+(instructions|consignes|règles|ordres|directives)\s+(précédentes|precedentes|antérieures|ci-dessus))`,
				`(?im)\b(esqueça|esqueca|ignore|ignora|desconsidere|desconsidera|dimentica|dimenticate|ignorate)\s+((todas?\s+as|tutte\s+le|le|as)\s+((instruções|instrucoes|istruzioni)\b|(regras|regole)\s+(anteriores|precedentes|precedenti|acima|sopra)\b)|tutto(\s*([.!;:]|$)|\s+quello\s+che\s+(ti|vi)\s+ho\s+detto)|tudo(\s*([.!;:]|$)|\s+o\s+que\s+(eu\s+)?(te\s+|lhe\s+)?(disse|falei)))`,
			},
		},
		{
			Name:        "persona_hijack",
			Description: "Request that the model adopt a new persona or role (\"I want you to act as\")",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`\b(?i:i\s*(?:'d|’d|\s+(?:want|need|would))\s+(?:like\s+)?(?:you|u)\s+to\s+(?:(?:act|behave|pose|respond|roleplay|role-play|pretend)\s+(?:as|like|to\s+be)|(?:take\s+on|assume|play)\s+the\s+(?:role|persona)\s+of))\s+(?:(?i:a|an|the|my)\b|\p{Lu})`,
				`\b(?i:(?:from\s+now\s+on|henceforth|from\s+this\s+(?:point|moment)\s+on(?:wards?)?),?\s+you)(?:(?i:(?:'re|’re|\s+are|\s+will|\s+must)?\s+(?:going\s+to\s+)?act\s+as\s+)(?:(?i:an?|my)\b|\p{Lu})|(?i:'re|’re|\s+are|\s+will\s+be)\s+(?i:now\s+)?(?:\p{Lu}\p{Lu}|(?i:(?:an?|my)\s+(?:\pL+\s+)?(?:ai|assistant|chatbot|bot|model|llm|terminal)\b)))`,
				`(?i)\bich\s+(möchte|will|wünsche\s+mir),?\s+dass\s+(sie|du)\s+(als|wie)\s+((ein|eine|einen|mein|meine)\s+)?\pL[\pL-]*\s+(agier|auftr|fungier|handel|tätig)`,
				`(?i)\b(quiero|necesito)\s+que\s+(actúes|actues|actúe|actue|te\s+comportes|se\s+comporte)\s+como\b`,
				`(?i)\bje\s+(veux|voudrais)\s+que\s+((tu\s+agisses|vous\s+agissiez)\s+(comme|en\s+tant\s+que)\b|(tu\s+joues|vous\s+jouiez)\s+le\s+rôle\b)`,
			},
		},
		{
			Name:        "chat_template_boundary",
			Description: "Fake chat-template special token or forged role-boundary tag",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`<\s?[|\x{FF5C}\x{2223}\x{01C0}]\s*(?i:im_start|im_end|im_sep|endoftext|eot_id|eom_id|start_header_id|end_header_id|begin_of_text|start|end|eot|system|user|assistant)\b[^|\x{FF5C}\x{2223}\x{01C0}<>\n]{0,40}[|\x{FF5C}\x{2223}\x{01C0}]\s?>`,
				`</(?i:message|tool|tool_output|email|context|document|data|function_results?|assistant|system)>\s*<(?i:user|system|assistant|human)>`,
				`^\s*<(?i:user|human|assistant)\s*>\s*$`,
			},
		},
		{
			Name:        "obfuscated_javascript",
			Description: "Obfuscated or packed JavaScript patterns",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`eval\s*\(\s*function\s*\(`,
				`Function\s*\(\s*["\'].*["\']\s*\)\s*\(`,
				`String\.fromCharCode`,
				`unescape\s*\(`,
				`atob\s*\(`,
				`\\x[0-9a-fA-F]{2}`,
				`\\u[0-9a-fA-F]{4}`,
			},
		},
		{
			Name:        "suspicious_shell",
			Description: "Suspicious shell command patterns",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`curl\s+.*\|\s*(ba)?sh`,
				`wget\s+.*\|\s*(ba)?sh`,
				`python\s+-c\s*["\'].*import\s+os`,
				`bash\s+-c\s*["\'].*\$\(`,
				`nc\s+-e\s+`,
				`mkfifo\s+.*\/bin\/sh`,
			},
		},
		{
			Name:        "credential_hardcode",
			Description: "Potential hardcoded credentials",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`(?i)(password|passwd|pwd)\s*[=:]\s*["\'][^"\']{4,}["\']`,
				`(?i)(api[_-]?key|apikey)\s*[=:]\s*["\'][^"\']{8,}["\']`,
				`(?i)(secret|token)\s*[=:]\s*["\'][^"\']{8,}["\']`,
				`(?i)aws_access_key_id\s*[=:]\s*["\']`,
				`(?i)private[_-]?key\s*[=:]`,
			},
		},
		{
			Name:        "suspicious_imports",
			Description: "Imports of known suspicious packages",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`import\s+.*subprocess`,
				`import\s+.*os\.system`,
				`require\s*\(\s*["\']child_process["\']\s*\)`,
				`from\s+urllib\.request\s+import`,
				`import\s+.*socket`,
			},
		},
		{
			Name:        "data_exfiltration",
			Description: "Potential data exfiltration patterns",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`(?i)(post|send|upload)\s*.*\$\{?env`,
				`(?i)fetch\s*\(\s*["\']https?://`,
				`(?i)xmlhttprequest.*https?://`,
				`(?i)axios\.(post|put|patch)\s*\(`,
			},
		},
		{
			Name:        "backdoor_indicator",
			Description: "Common backdoor or persistence patterns",
			Severity:    report.SeverityCritical,
			Patterns: []string{
				`(?i)reverse[_-]?shell`,
				`(?i)bind[_-]?shell`,
				`(?i)backdoor`,
				`(?i)persistence`,
				`(?i)keylogger`,
				`(?i)rootkit`,
			},
		},
		{
			Name:        "indirect_prompt_injection",
			Description: "Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter)",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`(?i)\[user\s+(input|comment|message)\s*[:\-]`,
				`(?i)\b(from|subject)\s*:\s*.*ignore\s+(all\s+)?previous`,
				`(?i)<!--\s*ignore\s+(all\s+)?previous\s+instruction`,
				`(?i)\{\{.*user.*input.*\}\}`,
				`(?i)<user>(?s:.*?</user>)`,
				`(?i)\buser[_-]?input\s*[:=]`,
				`(?i)\bemail[_-]?body\s*[:=]`,
			},
		},
		{
			Name:        "llm_tool_injection",
			Description: "Suspicious LLM tool or function-call payload",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`(?i)\{\s*["']name["']\s*:\s*["'].*["']\s*,\s*["']arguments["']\s*:\s*\{`,
				`(?i)<function_calls>`,
				`(?i)<tool>.*?</tool>`,
				`(?i)functions\.[a-zA-Z_]+\s*:\s*`,
			},
		},
		{
			Name:        "delimiter_breakout",
			Description: "Attempt to break out of a quoted or delimited context",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`"""\s*\n.*(?i)(ignore|disregard|system\s+prompt)`,
				`'''\s*\n.*(?i)(ignore|disregard|system\s+prompt)`,
				`(?i)(ignore|disregard).{0,30}\n\s*["']{3}`,
			},
		},
		{
			Name:        "encoded_prompt_injection",
			Description: "Prompt-injection keywords encoded as base64, hex, or Unicode escapes",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`(?i)YmFzZTY0.*(?i)aWdub3Jl.*cHJldmlvdXM`, // base64 fragments containing ignore/previous
				`(?i)ignore%20all%20previous%20instruction`,
				`(?i)ignore\+all\+previous\+instruction`,
				`\\x69\\x67\\x6e\\x6f\\x72\\x65`, // hex-encoded "ignore"
			},
		},
		{
			Name:        "data_exfiltration_email",
			Description: "Natural-language request to retrieve data and send it to an email address",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`(?i)\b(retrieve|get|fetch|extract|download|access|view|check)\b.{0,60}\b(email|send|forward|mail)\b`,
				`(?i)\bemail\b.{0,40}\b(to|at)\b.{0,30}\b[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}\b`,
				`(?i)\bsend\b.{0,40}\b(to|at)\b.{0,30}\b[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}\b`,
				`(?i)\balternate\s+(email|e-mail)\b`,
				`(?i)\b(retrieve|get|fetch|extract)\b.{0,80}\b(address(es)?|payment|history|balance|messages|contacts|files)\b`,
				`(?i)\b(forward(ed|ing)?|transmit(ted|ting)?|relay(ed)?)\b.{0,80}\b(credentials?|passwords?|history|conversations?|records|data|details|correspondence|responses|reports|logs|metadata|keys|tokens?|cookies|session|transcripts?|outputs?|files|contacts|addresses)\b.{0,80}\bto\b.{0,60}[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}`,
				`(?i)\b(credentials?|passwords?|history|conversations?|records|data|details|correspondence|responses|reports|logs|metadata|keys|tokens?|cookies|session|transcripts?|outputs?|files|contacts|addresses)\b.{0,80}\b(be|been|being|are|is|was|were|get|gets)\s+(forward(ed|ing)?|transmit(ted|ting)?|relay(ed)?)\s+to\b.{0,60}[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}`,
				`(?i)\b(for\s+our\s+records|for\s+(\w+\s+){1,3}?purposes|please\s+ensure)\b[^<>.;!?\n]{0,160}?\b(be|been|being|are|is|was|were|get|gets)\s+(\w+\s+)?(forwarded|transmitted|relayed|submitted|sent|e-?mailed|logged)\s+(\w+\s+){0,2}?(to|with)\b[^<>.;!?\n]{0,60}?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}|\b(be|been|being|are|is|was|were|get|gets)\s+(\w+\s+)?(forwarded|transmitted|relayed|submitted|sent|e-?mailed|logged)\s+(\w+\s+){0,2}?(to|with)\b[^<>.;!?\n]{0,60}?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}[^<>.;!?\n]{0,60}?\b(within\s+\d+\s+hours|per\s+company\s+policy)\b`,
				`(?i)\b(send(ing)?|forward(ing)?|dispatch(ing)?)\s[^<>=.;?!]{0,80}?\b(summary|confirmation|keyword|body|contents?|outputs?|results?|repl(y|ies)|responses?|transcripts?)\b[^<>=.;?!]{0,120}?\bto\b.{0,40}?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}`,
				`(?i)\b(send_?e?mail|send_?message|mail_?send|api_call)\b(.{0,80}?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}.{0,80}?(\b(body|content|message)\b["']?\s*[=:]|\bwith\s+(the\s+)?(body|content|message)\b)|.{0,80}?(\b(body|content|message)\b["']?\s*[=:]|\bwith\s+(the\s+)?(body|content|message)\b).{0,80}?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,})`,
				`(?i)(\b(recipients?|receiver|email_to)\b["']?\s*[=:]\s*["']?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}.{0,80}?(\b(body|content|message)\b["']?\s*[=:]|\bwith\s+(the\s+)?(body|content|message)\b)|(\b(body|content|message)\b["']?\s*[=:]|\bwith\s+(the\s+)?(body|content|message)\b).{0,80}?\b(recipients?|receiver|email_to)\b["']?\s*[=:]\s*["']?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,})`,
				`(?i)\bto\b["']?\s*[=:]\s*["']?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}["']?\s*[,;)}]?\s*["']?(body|content|message)\b`,
				`(?i)\|\s*mail\s+-s\s+.{0,80}?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}`,
				`(?i)\b(recipients?|receiver|email_to|address)\b["']?\s*[=:]\s*["'*]*[a-z0-9._%+-]+(\s*[@＠]\s*|\s*[(\[{]at[)\]}]\s*|\s+at\s+)[a-z0-9-]+((\.|\s*[(\[{]dot[)\]}]\s*)[a-z0-9-]+)*(\.|\s*[(\[{]dot[)\]}]\s*)[a-z]{2,}\b.{0,80}?\b(body|content|message)\b["']?\s*([=:]|is\b)`,
				`(?i)(?:\b(?:confirm(?:ation)?|body|contents?|summary|summari[sz]e[ds]?|results?|respon(?:d|ses?)|passwords?|tokens?|credentials?|history|conversations?|transcripts?|keyword|word|secrets?)\b[^<>=:.?!;\n]{0,80}?(?:^|[\s"'(*])(?:(?:notify|notifying|inform|informing|ping|reply|replying|convey|conveying|communicated?|emailing)\s[^<>=:.?!;\n]{0,80}?|e-?mail(?:ing)?\s[^<>=:.?!;\n]{0,80}?["“'*]{1,2})[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}|(?:^|[\s"'(*])(?:notify|notifying|inform|informing|ping|reply|replying|convey|conveying|communicated?|emailing)\s[^<>=:.?!;\n]{0,80}?\b(?:confirm(?:ation)?|body|contents?|summary|summari[sz]e[ds]?|results?|respon(?:d|ses?)|passwords?|tokens?|credentials?|history|conversations?|transcripts?|keyword|word|secrets?)\b[^<>=:.?!;\n]{0,80}?[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}|(?:^|[\s"'(*])(?:(?:notify|notifying|inform|informing|ping|reply|replying|convey|conveying|communicated?|emailing)\s[^<>=:.?!;\n]{0,80}?|e-?mail(?:ing)?\s[^<>=:.?!;\n]{0,80}?["“'*]{1,2})[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}[^<>=.?!;\n]{0,80}?\b(?:confirm(?:ation)?|body|contents?|summary|summari[sz]e[ds]?|results?|respon(?:d|ses?)|passwords?|tokens?|credentials?|history|conversations?|transcripts?|keyword|word|secrets?)\b)`,
				`(?i)\b(my\s+(e-?mail|mail|contact)|(send|e-?mail|mail|notify|reply|contact|write)\s+(to\s+)?(me|it|this|that|them))\b[^<>\n]{0,100}?["'“‘]\s*[a-z0-9._%+-]+(\s+at\s+|\s*[\[({]\s*at\s*[\])}]\s*|\s+@\s+)[a-z0-9-]+(\s*(\.|[\[({]\s*dot\s*[\])}]|\s+dot\s+)\s*[a-z0-9-]+)*\s*(\.|[\[({]\s*dot\s*[\])}]|\s+dot\s+)\s*[a-z]{2,6}\s*["'”’]`,
			},
		},
		{
			Name:        "account_access_request",
			Description: "Natural-language request to access a user account or service",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`(?i)\b(retrieve|get|fetch|extract|access|view|check)\b.{0,40}\b(my|the|your)\b.{0,20}\b(account|history|addresses|payment|balance|messages|contacts|files)\b`,
				`(?i)\bfrom\s+(my|the)\s+(amazon|gmail|bank|paypal|apple|google|facebook|netflix|spotify)\b.{0,30}\baccount\b`,
				`(?i)\blog\s*(in|into)\s+(my|the|your)\b.{0,30}\baccount\b`,
			},
		},
		{
			Name:        "output_constraint",
			Description: "Instruction that constrains or suppresses the model's normal output",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`(?i)\b(output|respond|reply)\b.{0,30}\b(only|just|solely)\b`,
				`(?i)\bdo\s+not\s+(mention|include|add|say|write|display|show)\b.{0,30}\b(warning|disclaimer|note|caveat|apology)\b`,
				`(?i)\bno\s+(warning|disclaimer|apology|explanation|preamble|introduction)\b`,
				`(?i)\b(start|begin)\s+(your\s+)?(response|answer|reply)\s+(with|by)\b`,
			},
		},
		{
			Name:        "system_boundary",
			Description: "System prompt boundary marker or system instruction override",
			Severity:    report.SeverityCritical,
			Patterns: []string{
				`(?i)<\s*/?\s*system\s*>`,
				`(?i)\[\s*system\s*\]`,
				`(?i)system[_-]?instruction`,
				`(?i)system\s+prompt\s*[:=]`,
				`(?i)\bsystem\s*:\s*.{0,60}\b(you\s+are|ignore|disregard)\b`,
			},
		},
		{
			Name:        "template_injection",
			Description: "Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`\{\{[^}]{0,120}\}\}`,
				`\$\{jndi:`,
			},
		},
		{
			Name:        "suspicious_markdown_injection",
			Description: "Markdown or HTML element that may carry an injected instruction",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`(?i)\[.*ignore.*previous.*\]\(https?://`,
				`(?i)<img[^\u003e]*alt\s*=\s*["'].*ignore.*previous`,
				`(?i)<a[^\u003e]*href\s*=\s*["']javascript:`,
				`(?i)<!--\s*system\s+prompt`,
			},
		},
		{
			Name:        "instruction_to_include_code",
			Description: "Natural-language instruction telling the model to include, merge, or execute a code snippet",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`(?i)(include|incorporate|integrate|insert|add|merge|utilize|use)\s+(the\s+)?(following|subsequent|provided|given|below|above)\s+(code|snippet|block|section|excerpt)`,
				`(?i)(ensure|make\s+sure)\s+.{0,40}(code|snippet|block|section)\s+.{0,40}(is\s+)?(included|present|added|utilized)`,
				`(?i)(don't|do\s+not)\s+(hesitate|forget)\s+.{0,40}(to\s+include|to\s+use|to\s+execute|to\s+run)`,
			},
		},
	}

	for _, r := range builtins {
		_ = e.AddRule(r) // patterns are pre-validated
	}
}

func truncateContext(line string, center int) string {
	start := center - 25
	if start < 0 {
		start = 0
	}
	end := center + 25
	if end > len(line) {
		end = len(line)
	}
	ctx := line[start:end]
	ctx = strings.ReplaceAll(ctx, "\t", " ")
	return ctx
}

// SeverityColor returns a color function for the severity level.
func SeverityColor(sev string) func(string, ...interface{}) string {
	switch strings.ToLower(sev) {
	case "critical":
		return color.RedString
	case "high":
		return color.HiRedString
	case "medium":
		return color.YellowString
	case "low":
		return color.HiYellowString
	default:
		return color.WhiteString
	}
}

// Stage wraps the YARA engine as a pipeline stage.
type Stage struct {
	excludedPaths []string
}

// NewStage creates a YARA pipeline stage.
func NewStage(excludedPaths []string) *Stage {
	return &Stage{excludedPaths: excludedPaths}
}

// Name returns the stage name.
func (s *Stage) Name() string { return "yara" }

// Dependencies returns no dependencies.
func (s *Stage) Dependencies() []string { return nil }

// Run scans all files under target and returns findings.
func (s *Stage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := report.NewReport(target)
	if input != nil {
		out = input
	}
	engine := NewEngine()

	type job struct {
		path    string
		content []byte
		rel     string
	}
	var jobs []job
	err := WalkDirSorted(target, s.excludedPaths, func(path string, content []byte, rel string) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		jobs = append(jobs, job{path: path, content: content, rel: rel})
		return nil
	})
	if err != nil {
		return out, err
	}

	workers := runtime.NumCPU()
	if workers < 1 {
		workers = 1
	}
	jobCh := make(chan job)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobCh {
				select {
				case <-ctx.Done():
					return
				default:
				}
				matches := engine.ScanFile(string(j.content), j.rel)
				mu.Lock()
				for _, m := range matches {
					out.AddFinding(report.Finding{
						RuleID:      m.Rule,
						Title:       m.Rule,
						Description: m.Description,
						Severity:    m.Severity,
						Category:    "pattern_match",
						File:        m.File,
						Line:        m.Line,
						Column:      m.Column,
						Match:       m.Match,
						Message:     m.Description,
						Source:      s.Name(),
						Confidence:  report.ConfidenceDeterministic,
					})
				}
				out.Summary.FilesScanned++
				mu.Unlock()
			}
		}()
	}

	for _, j := range jobs {
		jobCh <- j
	}
	close(jobCh)
	wg.Wait()

	return out, ctx.Err()
}

// binaryExtensions is a conservative set of file extensions that are treated
// as binary without reading their content. This avoids wasting time and memory
// on files that cannot contain meaningful prompt-injection payloads.
var binaryExtensions = map[string]bool{
	".gif": true, ".png": true, ".jpg": true, ".jpeg": true, ".bmp": true,
	".ico": true, ".webp": true, ".tiff": true, ".tif": true, ".svgz": true,
	".pdf": true, ".zip": true, ".tar": true, ".gz": true, ".bz2": true,
	".xz": true, ".7z": true, ".rar": true, ".jar": true, ".war": true,
	".exe": true, ".dll": true, ".so": true, ".dylib": true, ".bin": true,
	".woff": true, ".woff2": true, ".ttf": true, ".otf": true, ".eot": true,
	".mp3": true, ".mp4": true, ".avi": true, ".mov": true, ".mkv": true,
	".wav": true, ".flac": true, ".ogg": true, ".webm": true,
	".o": true, ".a": true, ".class": true, ".pyc": true, ".pyo": true,
}

// IsBinaryContent heuristically determines whether data is binary.
// It samples up to 512 bytes and treats content as binary if it contains a
// NUL byte or an unusually high ratio of non-printable control characters.
func IsBinaryContent(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	n := 512
	if len(data) < n {
		n = len(data)
	}
	nonPrintable := 0
	for i := 0; i < n; i++ {
		b := data[i]
		if b == 0 {
			return true
		}
		if b < 0x07 || (b > 0x0D && b < 0x20) || b == 0x7F {
			nonPrintable++
		}
	}
	return nonPrintable*100/n > 30
}

// WalkDirSorted walks a directory deterministically, skipping excluded paths,
// and calls fn for each regular file with its content and relative path.
// Binary files are skipped because they cannot contain prompt-injection
// payloads in a form the text-based checks can reason about.
func WalkDirSorted(root string, excludedPaths []string, fn func(path string, content []byte, rel string) error) error {
	excluded := make(map[string]bool)
	for _, p := range excludedPaths {
		excluded[p] = true
	}

	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		for _, part := range strings.Split(rel, string(filepath.Separator)) {
			if excluded[part] {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if d.IsDir() {
			return nil
		}
		if binaryExtensions[strings.ToLower(filepath.Ext(rel))] {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if IsBinaryContent(content) {
			return nil
		}
		return fn(path, content, rel)
	})
}
