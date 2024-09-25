package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(message)
	fmt.Println("Success create", name)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	message := ""
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(message)
	fmt.Println("Success add", name)
	return nil
}

func main() {
	createNewFile("test.txt", "File ini ditulis menggunakan Go os standard library")
	addToFile("test.txt", "Ini adalah tulisan yang ditambahkan melalui function addToFile")

	result, _ := readFile("test.txt")
	fmt.Println(result)
}
