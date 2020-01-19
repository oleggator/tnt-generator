package generator

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func Format(files ...string) error {
	args := append([]string{"-i"}, files...)

	cmd := exec.Command("/usr/local/bin/clang-format", args...)
	return cmd.Run()
}

func FormatStream(r io.Reader, w io.Writer) error {
	cmd := exec.Command(fmt.Sprintf("clang-format"))
	cmd.Stdout = w

	var b bytes.Buffer
	cmd.Stderr = &b

	if err := cmd.Run(); err != nil {
		if _, isExitError := err.(*exec.ExitError); isExitError {
			return fmt.Errorf("%s", b.String())
		}

		return err
	}

	return nil
}