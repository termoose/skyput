package main

import (
	"flag"
	_ "fmt"
	"github.com/fatih/color"
	"github.com/termoose/skyput/cache"
	"github.com/termoose/skyput/config"
	"github.com/termoose/skyput/portal"
	"github.com/termoose/skyput/upload"
	"os"
	"strconv"
)

func main() {
	portalSelector := flag.Bool("portal", false, "select portal")
	uploadList := flag.Bool("list", false, "list previous uploads")
	flag.Parse()

	c := config.Parse()

	flag.Usage = func() {
		c := color.New(color.FgGreen)
		c.Println("Usage: skynet filename [-portal] [-list n]")
		c.Println("\t-portal\tshow portal selector")
		c.Println("\t-list n\tlist n previous uploads")
	}

	if flag.NFlag() == 0 && flag.NArg() == 0 {
		flag.Usage()
		return
	}

	if *portalSelector {
		portal.Show(&c)
		return
	}

	if *uploadList {
		count := 10

		if flag.NArg() == 1 && flag.NFlag() == 1 {
			count, _ = strconv.Atoi(os.Args[2])
		}

		c, _ := cache.NewCache("cache")
		c.ShowLatest(count)
		return
	}

	selectedPortal := c.GetSelectedPortal()
	err := upload.Do(os.Args[1], selectedPortal)

	if err != nil {
		c := color.New(color.FgRed)
		c.Printf("%v\n", err)
	}
}
