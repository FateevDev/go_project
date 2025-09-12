package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		exitWithMessageAndErrorCode("You must provide only one file name", 1)
	}

	fileName := strings.TrimSpace(args[1])

	if fileName == "" {
		exitWithMessageAndErrorCode("Filename must not be empty", 1)
	}

	file, err := os.Open(fileName)

	if err != nil {
		exitWithMessageAndErrorCode("File not found: "+fileName, 1)
	}

	_, err = io.Copy(os.Stdout, file)

	if err != nil {
		exitWithMessageAndErrorCode(fmt.Sprintf("Error on print file content: %v", err), 1)
	}
}

func exitWithMessageAndErrorCode(m string, ec int) {
	fmt.Println(m)
	os.Exit(ec)
}
