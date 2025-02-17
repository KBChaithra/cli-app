package tasks

import (
	"log"
	"os"
)

func WriteFile(params map[string]string) {
	filename := params["filename"]
	content := params["content"]

	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		log.Printf("Error writing to file: %v", err)
	} else {
		log.Printf("File written successfully")
	}
}
