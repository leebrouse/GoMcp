package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("python3", "script.py", "arg1", "arg2")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run python script: %v", err)
	}
	fmt.Printf("%s", string(output))
}
