package main

import (
	"fmt"
	"os"

	"github.com/openware/goci/git"
	"github.com/openware/goci/versions"
	"github.com/openware/pkg/ika"
)

func actionVersions() error {
	// read configuration from the file and environment variables
	var cnf git.Config
	if err := ika.ReadConfig("", &cnf); err != nil {
		panic(err)
	}

	tmp := "./tmp"
	// Path to versions file
	if Path == "" {
		Path = fmt.Sprintf("opendax/%s/versions.yaml", cnf.Branch)
	}
	// Remove existing git folder
	if err := os.RemoveAll(tmp); err != nil {
		panic(err)
	}

	fmt.Printf("Clone the repository `%s`\n", cnf.Repo)
	fmt.Printf("Username: %s\n", cnf.Username)
	fmt.Printf("Email: %s\n", cnf.Email)

	auth := git.AuthToken{
		Username: cnf.Username,
		Token:    cnf.Token,
	}

	repo, err := git.Clone(&cnf, &auth, tmp)
	fmt.Println("Loading the versions file")
	v, err := versions.Load(fmt.Sprintf("%s/%s", tmp, Path))
	if err != nil {
		panic(err)
	}

	fmt.Println("Setting " + Component + " to " + Tag)
	v.SetTag(Component, Tag)

	fmt.Println("Saving the versions file")
	v.Save()

	// Commit & Push global OpenDAX versions
	fmt.Println("Commit & Push global OpenDAX versions")
	hash, err := git.Bump(repo, &auth, "Bump versions")
	if err == nil {
		fmt.Printf("Pushed with commit hash: %s", hash)
	}
	return err
}
