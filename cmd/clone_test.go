package cmd

import "testing"

func TestValidateGitURL(t *testing.T) {
	valid := []string{
		"https://github.com/user/repo.git",
		"http://github.com/user/repo.git",
		"ssh://git@github.com/user/repo.git",
		"git://github.com/user/repo.git",
		"git@github.com:user/repo.git",
	}
	for _, u := range valid {
		if err := validateGitURL(u); err != nil {
			t.Errorf("expected %q to be valid, got: %v", u, err)
		}
	}

	invalid := []string{
		"",
		"-u/tmp/pwn",
		"--upload-pack=bash -c 'evil'",
		"https://example.com/repo.git;rm -rf /",
		"git@github.com:user/repo.git|cat /etc/passwd",
		"https://example.com/repo.git\nBAD",
		"plain/path/without/scheme",
	}
	for _, u := range invalid {
		if err := validateGitURL(u); err == nil {
			t.Errorf("expected %q to be invalid", u)
		}
	}
}

func TestValidateCloneDir(t *testing.T) {
	valid := []string{"repo", "./repo", "tmp/repo-1", "/tmp/repo"}
	for _, d := range valid {
		if err := validateCloneDir(d); err != nil {
			t.Errorf("expected %q to be valid, got: %v", d, err)
		}
	}
	invalid := []string{"", "--option", "-u/tmp/pwn", "repo;rm -rf /", "repo\nBAD"}
	for _, d := range invalid {
		if err := validateCloneDir(d); err == nil {
			t.Errorf("expected %q to be invalid", d)
		}
	}
}
