package main

import (
	"flag"
	_ "fmt"
	"github.com/fatih/color"
	"github.com/termoose/skyput/cache"
	"github.com/termoose/skyput/config"
	"github.com/termoose/skyput/dappdappgo"
	"github.com/termoose/skyput/portal"
	"github.com/termoose/skyput/upload"
	"os"
	"strconv"
)

func main() {
	portalSelector := flag.Bool("portal", false, "select portal")
	uploadList := flag.Bool("list", false, "list previous uploads")
	ddg := flag.Bool("ddg", false, "make file searchable on dappdappgo")
	flag.Parse()

	c := config.Parse()

	flag.Usage = func() {
		c := color.New(color.FgGreen)
		c.Println("Usage: skynet [-portal] [-list n] filename")
		c.Println("\t-portal\tshow portal selector")
		c.Println("\t-list n\tshow the n previous uploads (default 10)")
		c.Println("\t-ddg\tmake the file searchable on dappdappgo")
	}

	if flag.NFlag() == 0 && flag.NArg() == 0 {
		flag.Usage()
		return
	}

	if *portalSelector {
		_, newDefault := portal.Show(c.GetPortals())
		c.SetDefaultPortal(newDefault)
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
	err, skyHash := upload.Do(os.Args[flag.NFlag() + flag.NArg()], selectedPortal)

	if *ddg {
		if flag.NArg() == 0 {
			c := color.New(color.FgYellow)
			c.Println("No filename specified!")
			return
		}

		err = dappdappgo.Post(skyHash)
	}

	if err != nil {
		c := color.New(color.FgRed)
		c.Printf("%v\n", err)
	}
}
