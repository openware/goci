package main

import (
	"fmt"

	"github.com/openware/goci/git"
	"github.com/openware/pkg/ika"
)

func actionClone() error {
	fmt.Println("Clone the repository")

	// read configuration from the file and environment variables
	var cnf git.Config
	if err := ika.ReadConfig("", &cnf); err != nil {
		panic(err)
	}

	auth := git.AuthToken{
		Username: cnf.Username,
		Token:    cnf.Token,
	}

	_, err := git.Clone(&cnf, &auth, "./tmp")
	return err
}
