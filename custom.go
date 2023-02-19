package main

import "fmt"

type Custom struct {
	Label         string  `json:"label"`
	CustomEntries []Entry `json:"entries"`
}

func (c Custom) Identifier() string {
	return "custom"
}

func (c Custom) defaultConfig() Custom {
	return Custom{
		Label:         "Custom",
		CustomEntries: []Entry{},
	}
}

func (c Custom) Entries() []Entry {
	entries := []Entry{}

	for _, entry := range c.CustomEntries {
		entries = append(entries, Entry{
			Name:     fmt.Sprintf("%s: %s", c.Label, entry.Name),
			Exec:     entry.Exec,
			Terminal: entry.Terminal,
		})
	}

	return entries
}
