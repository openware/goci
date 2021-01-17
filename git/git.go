package git

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Config for the git repository
type Config struct {
	Username string `env:"GIT_USERNAME" env-default:"kite-bot" env-description:"Git username"`
	Email    string `env:"GIT_EMAIL" env-default:"kite-bot@heliostech.fr" env-description:"Git user email"`
	Token    string `env:"GIT_TOKEN" env-description:"Git access token"`
	Repo     string `env:"GIT_REPO" env-default:"https://github.com/openware/versions.git" env-description:"Git repository url"`
	Branch   string `env:"DRONE_BRANCH" env-default:"2-6" env-description:"Drone target branch"`
}

// Auth to describe auth method
type Auth interface {
	Method() transport.AuthMethod
}

// AuthBasic for username, password auth
type AuthBasic struct {
	Username string
	Password string
}

// AuthToken for access token auth
type AuthToken struct {
	Username string
	Token    string
}

// Method implementation of AuthToken
func (a *AuthToken) Method() transport.AuthMethod {
	return &http.BasicAuth{
		Username: a.Username,
		Password: a.Token,
	}
}

// Clone repository with config
func Clone(cnf *Config, auth Auth, outDir string) (*git.Repository, error) {
	repo, err := git.PlainClone(outDir, false, &git.CloneOptions{
		Auth:     auth.Method(),
		URL:      cnf.Repo,
		Progress: os.Stdout,
	})
	cfg, err := repo.Config()
	if err != nil {
		return nil, err
	}
	cfg.User.Name = cnf.Username
	cfg.User.Email = cnf.Email
	repo.SetConfig(cfg)
	return repo, err
}

// Bump for git add, commit and push
func Bump(repo *git.Repository, auth Auth, msg string) (hash string, err error) {
	// worktree of the project using the go standard library
	w, err := repo.Worktree()
	if err != nil {
		return "", err
	}

	// Adds the new file to the staging area
	_, err = w.Add(".")
	if err != nil {
		return "", err
	}

	// Commits the current staging area to the repository, with the new file
	cmt, err := w.Commit(msg, &git.CommitOptions{})
	if err != nil {
		return "", err
	}

	// Push using default options
	return cmt.String(), repo.Push(&git.PushOptions{
		Auth: auth.Method(),
	})
}
