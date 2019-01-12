package utils

import (
	"os/exec"
	"strings"
)

func System(line_command string) ([]byte, error) {
    var command []string = strings.Split(line_command, " ")
    return exec.Command(command[0], command[1:]...).Output()
}
