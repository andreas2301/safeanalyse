package diff

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// ChangeType represents the type of change.
type ChangeType int

const (
	Equal ChangeType = iota
	Added
	Removed
)

// Change represents a single line difference.
type Change struct {
	Type ChangeType
	Line int      // line number in original (for Removed) or new (for Added)
	Text string
}

// Diff computes a line-by-line diff between two strings.
func Diff(original, modified string) []Change {
	origLines := strings.Split(original, "\n")
	modLines := strings.Split(modified, "\n")

	// Simple LCS-based diff
	return computeLCS(origLines, modLines)
}

// computeLCS performs a longest common subsequence diff.
func computeLCS(a, b []string) []Change {
	m, n := len(a), len(b)
	// Build LCS matrix
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	// Backtrack to build diff
	var changes []Change
	i, j := m, n
	for i > 0 || j > 0 {
		if i > 0 && j > 0 && a[i-1] == b[j-1] {
			changes = append(changes, Change{Type: Equal, Line: i, Text: a[i-1]})
			i--
			j--
		} else if j > 0 && (i == 0 || dp[i][j-1] >= dp[i-1][j]) {
			changes = append(changes, Change{Type: Added, Line: j, Text: b[j-1]})
			j--
		} else if i > 0 {
			changes = append(changes, Change{Type: Removed, Line: i, Text: a[i-1]})
			i--
		}
	}

	// Reverse
	for left, right := 0, len(changes)-1; left < right; left, right = left+1, right-1 {
		changes[left], changes[right] = changes[right], changes[left]
	}

	return changes
}

// Format returns a colored string representation of the diff.
func Format(changes []Change, context int) string {
	var b strings.Builder
	for _, c := range changes {
		switch c.Type {
		case Equal:
			if context > 0 {
				b.WriteString(fmt.Sprintf(" %3d | %s\n", c.Line, c.Text))
			}
		case Added:
			b.WriteString(color.GreenString("+%3d | %s\n", c.Line, c.Text))
		case Removed:
			b.WriteString(color.RedString("-%3d | %s\n", c.Line, c.Text))
		}
	}
	return b.String()
}

// Stats returns summary statistics.
func Stats(changes []Change) (added, removed, unchanged int) {
	for _, c := range changes {
		switch c.Type {
		case Added:
			added++
		case Removed:
			removed++
		case Equal:
			unchanged++
		}
	}
	return
}

// UnifiedDiff produces a unified diff format string.
func UnifiedDiff(changes []Change, origLabel, modLabel string) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("--- %s\n", origLabel))
	b.WriteString(fmt.Sprintf("+++ %s\n", modLabel))

	i := 0
	for i < len(changes) {
		// Find a hunk
		if changes[i].Type == Equal {
			i++
			continue
		}

		// Find hunk boundaries
		start := i
		end := i
		for end < len(changes) && changes[end].Type != Equal {
			end++
		}
		// Include some context
		ctxStart := start - 3
		if ctxStart < 0 {
			ctxStart = 0
		}
		ctxEnd := end + 3
		if ctxEnd > len(changes) {
			ctxEnd = len(changes)
		}

		// Count lines
		origStart, origCount := 0, 0
		modStart, modCount := 0, 0
		for j := ctxStart; j < ctxEnd; j++ {
			if changes[j].Type != Added {
				if origCount == 0 {
					origStart = changes[j].Line
				}
				origCount++
			}
			if changes[j].Type != Removed {
				if modCount == 0 {
					modStart = changes[j].Line
				}
				modCount++
			}
		}

		b.WriteString(fmt.Sprintf("@@ -%d,%d +%d,%d @@\n", origStart, origCount, modStart, modCount))
		for j := ctxStart; j < ctxEnd; j++ {
			switch changes[j].Type {
			case Equal:
				b.WriteString(" " + changes[j].Text + "\n")
			case Added:
				b.WriteString("+" + changes[j].Text + "\n")
			case Removed:
				b.WriteString("-" + changes[j].Text + "\n")
			}
		}

		i = ctxEnd
	}
	return b.String()
}
