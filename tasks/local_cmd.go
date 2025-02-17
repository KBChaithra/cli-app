package tasks

import (
	"log"
	"os/exec"
)

func RunLocalCmd(command string) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing command: %v", err)
	} else {
		log.Printf("Command output: %s", output)
	}
}
