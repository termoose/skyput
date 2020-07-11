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
	all := flag.Bool("all", false, "upload all files in current directory")
	flag.Parse()

	c := config.Parse()

	flag.Usage = func() {
		c := color.New(color.FgGreen)

		c.Println("Usage: skynet [-portal] [-list n] filename")

		flag.VisitAll(func(f *flag.Flag) {
			c.Printf("\t-%s\t%s\n", f.Name, f.Usage)
		})
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

	var uploadString string
	if *all {
		uploadString = "*"
	} else {
		uploadString = os.Args[flag.NFlag()+flag.NArg()]
	}

	err, skyHashes := upload.Do(uploadString, selectedPortal)

	if *ddg {
		if flag.NArg() == 0 {
			c := color.New(color.FgYellow)
			c.Println("No filename specified!")
			return
		}

		if err == nil {
			for _, h := range skyHashes {
				err = dappdappgo.Post(h)
			}
		}
	}

	if err != nil {
		c := color.New(color.FgRed)
		c.Printf("%v\n", err)
	}
}
