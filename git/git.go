package git

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Config for the git repository
type Config struct {
	Username string `env:"GIT_USERNAME" env-description:"Git username"`
	Email    string `env:"GIT_EMAIL" env-description:"Git user email"`
	Token    string `env:"GIT_TOKEN" env-description:"Git access token"`
	Repo     string `env:"GIT_REPO" env-description:"Git repository url"`
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
	return git.PlainClone(outDir, false, &git.CloneOptions{
		Auth:     auth.Method(),
		URL:      cnf.Repo,
		Progress: os.Stdout,
	})
}
