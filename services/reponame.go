package services

import (
	"context"
	"findkey/github"
	"fmt"
	"net/http"
)

func GetList(session *http.Client, login string) []*github.Repository {
	client := github.NewClient(session)
	ctx := context.Background()

	opt := &github.RepositoryListOptions{
		Type:    "sources",
		Perpage: "100000",
	}

	// looping in total repo
	// 30 row perpage

	//var allRepos []*github.Repository
	perpage := "100000"
	repos, _, err := client.Repositories.List(ctx, login, perpage, opt)
	if err != nil {
		fmt.Println("errr")
	}

	var allRepos []*github.Repository

	loop := 0
	// Looping
	for _, repo := range repos {
		if *repo.Fork == false {
			loop = loop + 1
			// allRepos = append(allRepos, repos.Name)
			allRepos = append(allRepos, repos...)
		}
	}

	return allRepos
}
