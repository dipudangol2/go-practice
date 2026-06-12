package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File Handling in go:")
	fileName := "demo.txt"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Println("File not found... \nCreating file \"", fileName, "\"")
		file, fileErr := os.Create(fileName)
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}
		defer file.Close()
	} else {
		fmt.Println("File Exists!")
	}
}
