package common

import (
	"fmt"
	"os/exec"
)

func RunCommand(command, path string) (err error) {
	cmd := exec.Command(command)
	cmd.Dir = path
	fmt.Println(cmd.Dir)
	fmt.Println(command)
	fmt.Println(cmd.Args)
	fmt.Println(cmd)
	err = cmd.Run()
	return
}
