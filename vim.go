package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Vim struct {
	SessionDir string `json:"session_dir"`
	Editor     string `json:"editor"`
	Label      string `json:"label"`
}

func (v Vim) Identifier() string {
	return "vim"
}

func (v Vim) defaultConfig() Vim {
	return Vim{
		SessionDir: "",
		Editor:     "vim",
		Label:      "Vim",
	}
}

func (v Vim) IsAvailable(config Config) bool {
	return config.containsModule(v.Identifier()) && v.SessionDir != "" && config.Terminal != ""
}

func (v Vim) Entries() []Entry {
	entries := []Entry{}

	files, err := ioutil.ReadDir(v.SessionDir)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		entries = append(entries, Entry{
			Name:     fmt.Sprintf("%s: %s", v.Label, file.Name()),
			Terminal: true,
			Exec:     fmt.Sprintf("%s -S %s", v.Editor, filepath.Join(v.SessionDir, file.Name())),
		})
	}

	return entries
}
