package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Vim struct{}

func (v Vim) Identifier() string {
	return "vim"
}

func (v Vim) IsAvailable(config Config) bool {
	return config.containsModule(v.Identifier()) && config.VimConfig.SessionDir != "" && config.Terminal != ""
}

func (v Vim) Entries(config Config) []Entry {
	entries := []Entry{}

	files, err := ioutil.ReadDir(config.VimConfig.SessionDir)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		entries = append(entries, Entry{
			Name:     fmt.Sprintf("%s: %s", config.VimConfig.Label, file.Name()),
			Terminal: true,
			Exec:     fmt.Sprintf("%s -S %s", config.VimConfig.Editor, filepath.Join(config.VimConfig.SessionDir, file.Name())),
		})
	}

	return entries
}
