package go_embed

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

//go:embed goland.png
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("logo_new.png", logo, 0o644)
	if err != nil {
		log.Fatal(err)
	}
}
