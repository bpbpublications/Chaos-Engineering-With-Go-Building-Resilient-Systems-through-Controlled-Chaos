package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//  path where the dummy files will be created
	directoryPath := "/tmp/fault_injection"
	// Total disk space to be consumed (in MB)
	totalDiskSpace := 100
	// File size of each dummy file (in MB)
	fileSize := 10
	// Calculate the number of files needed to consume the total disk space
	noFiles := totalDiskSpace / fileSize
	// Create the directory if it doesn't exist
	err := os.MkdirAll(directoryPath, 0755)
	if err != nil {
		fmt.Println("Error creating the directory:", err)
		return
	}
	// Generate and write dummy files to consume disk space
	for i := 0; i < noFiles; i++ {
		fileName := fmt.Sprintf("file_%d.txt", i+1)
		filePath := fmt.Sprintf("%s/%s", directoryPath, fileName)
		err := createDummyFile(filePath, fileSize)
		if err != nil {
			fmt.Println("Error creating dummy file:", err)
			return
		}
	}
	fmt.Printf("%d MB disk space consumed in %s\n", totalDiskSpace, directoryPath)
}

// Function to write a dummy file of the given size (in MB)
func createDummyFile(filePath string, fileSize int) error {
	content := make([]byte, fileSize*1024*1024)
	return ioutil.WriteFile(filePath, content, 0644)
}
