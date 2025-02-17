package tasks

import (
	"fmt"
	"log"
	"os/exec"
)

func RunSSHCmd(params map[string]string) {
	user := params["user"]
	host := params["host"]
	command := params["command"]

	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", user, host), command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing SSH command: %v", err)
	} else {
		log.Printf("SSH command output: %s", output)
	}
}
