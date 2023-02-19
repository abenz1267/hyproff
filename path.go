package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Path struct{}

func (p Path) Identifier() string {
	return "path"
}

func (p Path) Entries() []Entry {
	entries := []Entry{}

	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")

	for _, v := range paths {
		files, err := ioutil.ReadDir(v)
		if err != nil {
			log.Println(err)
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			entries = append(entries, Entry{
				Name: fmt.Sprintf("Path: %s", file.Name()),
				Exec: file.Name(),
			})
		}
	}

	return entries
}
