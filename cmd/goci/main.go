package main

import (
	"fmt"
	"os"

	"github.com/openware/pkg/kli"
)

// Version of the command line
var Version = "SNAPSHOT"

// Path to versions file
var Path string

// Component name
var Component string

// Tag of the component
var Tag string

// Changed file display depth(dir1/dir2/.../dir*depth*)
var Depth = 0

func main() {
	cli := kli.NewCli("goci", "Openware versions cli", Version)

	cmdVersions := kli.NewCommand("versions", "Update component version in openware/versions").Action(actionVersions)
	cli.StringFlag("path", "Path to the versions folder (e.g. opendax/2-6/versions.yaml)", &Path)
	cli.StringFlag("component", "Name of the component to update", &Component)
	cli.StringFlag("tag", "Tag to insert into the versions file", &Tag)
	cli.AddCommand(cmdVersions)

	cmdChanges := kli.NewCommand("changes", "List files changed in the last commit").Action(actionChanges)
	cli.IntFlag("depth", "Depth of directories changed in the latest git commit", &Depth)
	cli.AddCommand(cmdChanges)

	if err := cli.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
		os.Exit(1)
	}
}
