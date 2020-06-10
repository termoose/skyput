package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/termoose/skyput/upload"
	"os"
)

func main() {
	portal := flag.Bool("portal", false, "select portal")
	flag.Parse()

	flag.Usage = func() {
		c := color.New(color.FgGreen)
		c.Println("Usage: skynet filename [-portal]")
		c.Println("\t-portal\tshow portal selector")
	}

	if flag.NFlag() == 0 && flag.NArg() == 0 {
		flag.Usage()
		return
	}

	if *portal {
		fmt.Printf("Select portal? %v\n", *portal)
		return
	}

	err := upload.Do(os.Args[1])

	if err != nil {
		c := color.New(color.FgRed)
		c.Printf("%v\n", err)
	}
}
