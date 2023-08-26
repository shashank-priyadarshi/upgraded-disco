package common

import (
	"fmt"
	"os/exec"

	logger "github.com/rs/zerolog/log"
)

func RunCommand(command, path string) (err error) {
	cmd := exec.Command(command, path)
	// cmd.Dir = path
	// err = cmd.Run()
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Info().Msg(fmt.Sprint(err) + ": " + string(output))
		return
	}
	return
}
