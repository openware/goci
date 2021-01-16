package main

import (
	"fmt"
	"os"

	"github.com/openware/pkg/kli"
)

// Version of the command line
var Version = "SNAPSHOT"

// Path to versions file
var Path = "opendax/2-6/versions.yaml"

// Component name
var Component = "peatio"

// Tag of the component
var Tag string

func main() {
	cli := kli.NewCli("goci", "Openware versions cli", Version)

	cmdVersions := kli.NewCommand("versions", "Update component version in openware/versions").Action(actionVersions)
	cli.StringFlag("path", "Path to the versions folder (e.g. opendax/2-6/versions.yaml)", &Path)
	cli.StringFlag("component", "Name of the component to update", &Component)
	cli.StringFlag("tag", "Tag to insert into the versions file", &Tag)
	cli.AddCommand(cmdVersions)

	if err := cli.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
		os.Exit(1)
	}
}
