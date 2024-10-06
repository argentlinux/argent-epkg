package search

import (
	"bytes"
	"fmt"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"os/exec"
)

func RunSearch(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("command failed: %w", err)
	}
	return string(output), nil
}

func EmergeInfo(packageBuffer string) {
	var auth *widget.Button
	auth = widget.NewButton("Authorize", func() {
		auth.Disable()
		cmd := exec.Command("emerge --info -v", packageBuffer)

		buffer := bytes.Buffer{}
		buffer.Write([]byte("\n"))
		cmd.Stdin = &buffer

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err3 := cmd.Run()

		if err3 != nil {
			log.Println("Simple err", err3)
		} else {
		}
		auth.Enable()
	})
}
