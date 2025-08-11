package helper

import (
	"fmt"
	"os"
)

// helper function to Parse MarkDown to string
func ParseMarkDown(path, reference string) (string, error) {
	tmplBytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintln("Fail to read templete:", err), err
	}

	// create a new prompt fot generating text 
	newPrompt := fmt.Sprintf(string(tmplBytes), reference)

	return newPrompt, nil
}

