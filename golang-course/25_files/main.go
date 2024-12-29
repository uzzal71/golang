package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		// log the error
		panic(err)
	}

	fmt.Println("file name:", fileInfo.Name())
	fmt.Println("file or folder:", fileInfo.IsDir())
	fmt.Println("file size:", fileInfo.Size())
	fmt.Println("file permission:", fileInfo.Mode())
	fmt.Println("file modified at:", fileInfo.ModTime())
}
