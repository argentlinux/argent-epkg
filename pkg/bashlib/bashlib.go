package bashlib

import (
	"os/exec"
)

// most basic test
func BashLs() string {
	return "ls -la"
}

// grep test with pattern
// not used
func BashGrep(pattern, file string) string {
	return "grep '" + pattern + "' " + file
}

func RunBashCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func RunGentooQuery() (string, error) {
	command := "qlist -Iv"
	return RunBashCommand(command)
}
