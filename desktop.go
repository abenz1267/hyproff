package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/rkoesters/xdg/desktop"
)

var desktopFileLocations = []string{"/usr/share/applications/", "/usr/local/share/applications/", "~/.local/share/applications/"}

type Desktop struct{}

func (d *Desktop) Entries() []Entry {
	entries := []Entry{}

	for _, location := range desktopFileLocations {
		files, err := ioutil.ReadDir(location)
		if err != nil {
			log.Println(err)

			continue
		}

		for _, file := range files {
			if !file.IsDir() {
				r, err := os.Open(location + file.Name())
				if err != nil {
					log.Println(err)

					continue
				}

				d, err := desktop.New(r)
				if err != nil {
					log.Println(err)

					continue
				}

				if d.Type == desktop.Application && !d.NoDisplay {
					for _, action := range d.Actions {
						entries = append(entries, Entry{
							Name: fmt.Sprintf("Desktop: %s - %s", d.Name, action.Name),
							Exec: strings.ReplaceAll(action.Exec, "\"", "'"),
						})
					}
				}
			}
		}
	}

	return entries
}
