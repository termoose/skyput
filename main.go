package main

import (
	"github.com/fatih/color"
	"github.com/termoose/skyput/upload"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		c := color.New(color.FgYellow)
		c.Printf("🏥 Usage: %s [filename]\n", os.Args[0])
		return
	}

	err := upload.Do(os.Args[1])

	if err != nil {
		c := color.New(color.FgRed)
		c.Printf("%v\n", err)
	}
}
