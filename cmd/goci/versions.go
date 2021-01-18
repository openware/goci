package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	fmt.Printf("Cloning the repository `%s`\n", cnf.URL)
	fmt.Printf("Username: %s\n", cnf.Username)
	fmt.Printf("Email: %s\n", cnf.Email)

	auth := git.AuthToken{
		Username: cnf.Username,
		Token:    cnf.Token,
	}

	repo, err := git.Clone(&cnf, &auth, tmp)
	if err != nil {
		panic(err)
	}

	fmt.Println("Loading the versions file")
	v, err := versions.Load(fmt.Sprintf("%s/%s", tmp, Path))
	if err != nil {
		panic(err)
	}

	if Tag == "" {
		// Read .tags if exists to get tag version
		Tag, err = getTag()
		if err != nil {
			panic(errors.New("Tag is missing"))
		}
	}

	if Component == "" {
		Component = cnf.Repo
	}

	fmt.Println("Setting " + Component + " to " + Tag)
	v.SetTag(Component, Tag)

	fmt.Println("Saving the versions file")
	v.Save()

	// Commit & Push to global OpenDAX versions
	fmt.Println("Committing & Pushing to global OpenDAX versions")
	hash, err := git.Update(repo, &auth, fmt.Sprintf("%s: Update %s version to %s", cnf.Branch, Component, Tag))
	if err == nil {
		fmt.Printf("Pushed with commit hash: %s\n", hash)
	}
	return err
}

func getTag() (string, error) {
	val, err := ioutil.ReadFile(".tags")
	if err == nil {
		return strings.TrimSuffix(string(val), "\n"), nil
	}
	return "", err
}
