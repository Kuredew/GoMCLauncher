package utils

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func ExecuteCMD(firstArg string, secondArg ...string) error {
	cmd := exec.Command(firstArg, secondArg...)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Printf("Error calling java client : %s", err)
		return err
	}

	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	return cmd.Wait()
}