package main

import (
	"embed"
	"fmt"
	"log"
	"os"
)

//go:embed test/version.txt
var version string

//go:embed test/goland.png
var logo []byte

//go:embed test/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("logo_new.png", logo, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	dirEntries, _ := path.ReadDir("test/files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("test/files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
