package common

import (
	"fmt"
	"os/exec"
)

func RunCommand(command, path string) (err error) {
	cmd := exec.Command(command, path)
	// cmd.Dir = path
	// err = cmd.Run()
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	return
}
