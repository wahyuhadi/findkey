package services

import (
	"context"
	"findkey/github"
	"fmt"
	"net/http"
)

// GetListCommitsSha function to get sha for all commits in user repos
// example : GET https://api.github.com/repos/wahyuhadi/python-shodan-cli/commits?per_page=100000
func GetListCommitsSha(session *http.Client, user string, repos []string, perpage string) {
	client := github.NewClient(session)
	ctx := context.Background()

	for _, reponame := range repos {
		// ListCommitsSha function in github/repos_commits.go
		commits, _, err := client.Repositories.ListCommitsSha(ctx, user, reponame, perpage, nil)

		if err != nil {
			fmt.Println(err)
		}

		for _, sha := range commits {
			fmt.Println(reponame, *sha.SHA)
		}
	}
}
