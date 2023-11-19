package main

import (
	"fmt"
	"io"
	"os"
)

const SkyBlue = "\033[1;36m"

func exit(errMsg string) {
	fmt.Print(errMsg)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		exit("Usage: gat <file>")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		exit(err.Error())
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		exit(err.Error())
	}
	if fileInfo.IsDir() {
		errMsg := fmt.Sprintf("error: %s is a directory\n", os.Args[1])
		exit(errMsg)
	}

	data := make([]byte, fileInfo.Size())
	for {
		n, err := f.Read(data)
		fmt.Print(SkyBlue)
		os.Stdout.Write(data[:n])
		if err != nil {
			if err != io.EOF {
				exit(err.Error())
			}
			break
		}
	}
}
