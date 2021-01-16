package main

import (
	"fmt"

	"github.com/openware/goci/git"
	"github.com/openware/pkg/ika"
)

func actionClone() error {
	// read configuration from the file and environment variables
	var cnf git.Config
	if err := ika.ReadConfig("", &cnf); err != nil {
		panic(err)
	}

	fmt.Printf("Clone the repository `%s`\n", cnf.Repo)
	fmt.Printf("Username: %s\n", cnf.Username)
	fmt.Printf("Email: %s\n", cnf.Email)

	auth := git.AuthToken{
		Username: cnf.Username,
		Token:    cnf.Token,
	}

	_, err := git.Clone(&cnf, &auth, "./tmp")
	return err
}
