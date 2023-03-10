package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	PermFile   = 0o644
	PermFolder = 0o755
)

type Entry struct {
	Name     string
	Exec     string
	Terminal bool
}

type Module interface {
	Entries() []Entry
	Identifier() string
}

type Config struct {
	Terminal string   `json:"terminal"`
	Modules  []string `json:"modules"`
	Vim      Vim      `json:"vim"`
	Custom   Custom   `json:"custom"`
}

func (c Config) containsModule(module string) bool {
	for _, v := range c.Modules {
		if v == module {
			return true
		}
	}

	return false
}

func main() {
	config := loadConfig()
	enabled := []Module{}

	modules := []Module{Hyprland{}, Path{}, Desktop{}, config.Vim, config.Custom}

	for _, v := range config.Modules {
		for _, module := range modules {
			if module.Identifier() == v {
				enabled = append(enabled, module)
			}
		}
	}

	for k, collector := range enabled {
		fmt.Printf("%%base_score=%d\n", k)

		for _, entry := range collector.Entries() {
			if entry.Terminal {
				fmt.Printf("%s=%s %s\n", entry.Name, config.Terminal, entry.Exec)

				continue
			}

			fmt.Printf("%s=%s\n", entry.Name, entry.Exec)
		}
	}
}

func loadConfig() Config {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	config := filepath.Join(configDir, "hyproff", "config.json")
	if _, err := os.Stat(config); err != nil {
		return createDefaultConfig(configDir)
	}

	var c Config
	b, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func createDefaultConfig(configDir string) Config {
	dir := filepath.Join(configDir, "hyproff")

	c := Config{
		Modules: []string{Custom{}.Identifier(), Hyprland{}.Identifier(), Vim{}.Identifier(), Desktop{}.Identifier(), Path{}.Identifier()},
		Vim:     Vim{}.defaultConfig(),
		Custom:  Custom{}.defaultConfig(),
	}

	err := os.MkdirAll(dir, PermFolder)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(&c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(dir, "config.json"), b, PermFile)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
