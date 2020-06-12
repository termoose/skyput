package main

import (
	"flag"
	_ "fmt"
	"github.com/fatih/color"
	"github.com/termoose/skyput/config"
	"github.com/termoose/skyput/portal"
	"github.com/termoose/skyput/upload"
	"os"
)

func main() {
	portalSelector := flag.Bool("portal", false, "select portal")
	flag.Parse()

	c := config.Parse()

	flag.Usage = func() {
		c := color.New(color.FgGreen)
		c.Println("Usage: skynet filename [-portal]")
		c.Println("\t-portal\tshow portal selector")
	}

	if flag.NFlag() == 0 && flag.NArg() == 0 {
		flag.Usage()
		return
	}

	if *portalSelector {
		portal.Show(&c)
		return
	}

	selectedPortal := c.GetSelectedPortal()
	err := upload.Do(os.Args[1], selectedPortal)

	if err != nil {
		c := color.New(color.FgRed)
		c.Printf("%v\n", err)
	}
}
