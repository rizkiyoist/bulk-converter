package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const targetConvert = ".mp3"

func main() {
	// find all filenames in this directory
	fileNames, err := findFileNames()
	if err != nil {
		fmt.Println(err)
	}
	for _, fileName := range fileNames {
		fmt.Println(fileName)
		noExtFileName := removeExtension(fileName)
		err = exec.Command("ffmpeg", "-n", "-i", fileName, "-vn", "-q:a", "0", noExtFileName+targetConvert).Run()
		fmt.Println(err)
	}
	fmt.Println("done")
}

func findFileNames() ([]string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed reading files: %v", err)
	}
	fileNames := []string{}

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

func removeExtension(fileName string) string {
	dot := strings.LastIndexByte(fileName, '.')
	if dot != -1 {
		fileName = fileName[:dot]
	}
	return fileName
}
