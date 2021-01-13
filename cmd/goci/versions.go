package main

import (
	"fmt"

	"github.com/openware/goci/versions"
)

func actionVersions() error {
	fmt.Println("Loading the versions file")
	v, err := versions.Load(Path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Setting " + Component + " to " + Tag)
	v.SetTag(Component, Tag)

	fmt.Println("Saving the versions file")
	return v.Save()
}
