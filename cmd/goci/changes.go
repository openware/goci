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

	var changedFiles []string

	for _, stat := range stats {
		changedFiles = append(changedFiles, stat.Name)
	}

	if Depth == 0 {
		fmt.Printf("%s\n", strings.Join(changedFiles, " "))
	} else {
		var res []string
		keys := make(map[string]int)

		for _, f := range changedFiles {
			if _, ok := keys[f]; !ok {
				keys[f] = 1
				filepath := strings.Split(f, "/")

				if len(filepath) >= Depth {
					filepath = filepath[0:Depth]
				}

				res = append(res, strings.Join(filepath, "/"))
			}
		}

		fmt.Printf("%s\n", strings.Join(res, " "))
	}

	return nil
}
