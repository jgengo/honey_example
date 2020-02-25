package main

import (
	"log"
	"os/exec"
)

// RepoCreate ...
type RepoCreate struct {
	URL string `json:"url"`
	Key string `json:"key"`
}

// GitClone ...
func GitClone(url, name string) error {
	cmd := exec.Command("git", "clone", url, "/tmp/"+name)
	log.Printf("cloning: (processing) %s\n", cmd.String())
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
