package main

import (
	"fmt"
	"os"
)

/*
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
*/

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	buf := make([]byte, 10)
	d, err := file.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(d, buf)

}
