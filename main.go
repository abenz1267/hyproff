package main

import (
	"fmt"
)

type Entry struct {
	Name string
	Exec string
}

type Collector interface {
	Entries() []Entry
}

func main() {
	collectors := []Collector{}

	hyprland := &Hyprland{}
	path := &Path{}

	collectors = append(collectors, hyprland, path)

	for _, collector := range collectors {
		for _, entry := range collector.Entries() {
			fmt.Printf("%s=%s\n", entry.Name, entry.Exec)
		}
	}
}