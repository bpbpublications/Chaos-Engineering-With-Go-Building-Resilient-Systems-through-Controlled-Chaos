package main

import (
	"fmt"
	"os/exec"
)

func main() {
	processName := "chrome.exe"

	err := killProcessByName(processName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Process terminated successfully.")
}

func killProcessByName(name string) error {
	// Execute the taskkill command to terminate processes by name
	cmd := exec.Command("taskkill", "/F", "/IM", name)
	err := cmd.Run()
	if err != nil {
		// Check if the error is due to no processes being found
		if exitErr, ok := err.(*exec.ExitError); ok {
			// If the exit status is 128, it indicates no processes found
			if exitErr.ExitCode() == 128 {
				return fmt.Errorf("no processes found matching the name '%s'", name)
			}
		}
		return err
	}

	return nil
}
