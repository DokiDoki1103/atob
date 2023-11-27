package gitx

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"os"
	"strings"
)

func Clone(url, path string) error {
	o := &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	}

	if strings.HasPrefix(url, "git@github.com") {
		var auth *ssh.PublicKeys
		sshKey, err := os.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
		if err != nil {
			return err
		}
		auth, keyError := ssh.NewPublicKeys("git", sshKey, "")
		if keyError != nil {
			return err
		}
		o.Auth = auth
	}

	repo, err := git.PlainClone(path, false, o)
	if err != nil {

		if err.Error() == "repository already exists" {
			repo, err := git.PlainOpen(path)
			if err != nil {
				return err
			}
			w, err := repo.Worktree()
			err = w.Pull(&git.PullOptions{
				Auth: o.Auth,
			})
			if err != nil {
				return err
			}
			fmt.Println(w.Status())

		}

		return err
	}

	log, err := repo.Log(&git.LogOptions{})
	if err != nil {
		return err
	}
	next, err := log.Next()
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(next)
	if err != nil {
		return err
	}
	fmt.Println(string(marshal))
	return nil
}
