package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Hyprland struct{}

func (h Hyprland) Identifier() string {
	return "hyprland"
}

func (h Hyprland) Entries() []Entry {
	cmd := exec.Command("hyprctl", "clients")

	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}

	entries := []Entry{}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(scanner.Text(), "Window ") {
			entries = append(entries, Entry{})
			continue
		}

		if strings.HasPrefix(line, "title:") {
			entries[len(entries)-1].Name = fmt.Sprintf("Window: %s", strings.TrimPrefix(line, "title: "))
			continue
		}

		if strings.HasPrefix(line, "pid:") {
			entries[len(entries)-1].Exec = fmt.Sprintf("hyprctl dispatch focuswindow pid:%s", strings.TrimPrefix(line, "pid: "))
			continue
		}

	}

	return entries
}
