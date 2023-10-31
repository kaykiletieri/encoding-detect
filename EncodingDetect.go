package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func main() {
	fmt.Print("Enter the file name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	foundEncoding, _, _ := charset.DetermineEncoding(buffer, "")
	if foundEncoding == nil {
		fmt.Println("Encoding not detected")
	} else {
		fmt.Println("Detected encoding:", foundEncoding)
	}
}
