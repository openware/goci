package main

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
)

func actionChanges() error {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	ref, err := repo.Head()
	if err != nil {
		return err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	stats, err := commit.Stats()
	if err != nil {
		return err
	}

	var paths []string

	for _, stat := range stats {
		paths = append(paths, stat.Name)
	}

	if Depth > 0 {
		paths = filterUnique(paths, Depth)
	}
	fmt.Printf("%s\n", strings.Join(paths, " "))

	return nil
}

func filterUnique(changes []string, depth int) []string {
	var paths []string
	keys := make(map[string]int)

	for _, f := range changes {
		filepath := strings.Split(f, "/")
		if len(filepath) >= depth {
			filepath = filepath[0:depth]
		}

		path := strings.Join(filepath, "/")
		if _, ok := keys[path]; !ok {
			keys[path] = 1
			paths = append(paths, path)
		}
	}
	return paths
}
