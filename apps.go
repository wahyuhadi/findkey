package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	user = flag.String("user", "google", "username target")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("github")},
	)
	tc := oauth2.NewClient(ctx, ts)

	get(tc, *user)
}

func get(session *http.Client, login string) {
	client := github.NewClient(session)
	ctx := context.Background()

	//opt := &github.RepositoryListAllOptions{Sort: "updated", Direction: "desc"}

	opt := &github.RepositoryListOptions{
		Type: "sources",
	}

	// list all repositories for the authenticated user
	for {
		repos, resp, err := client.Repositories.List(ctx, login, opt)
		if err != nil {
			fmt.Println("errr")
		}

		loop := 0
		for _, repo := range repos {
			loop = loop + 1
			fmt.Println(*repo.HTMLURL)
			fmt.Println(*repo.Name)
			fmt.Println(*repo.Fork)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

}
