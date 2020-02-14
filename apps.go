package main

import (
	"context"
	service "findkey/services"
	"flag"
	"os"

	"golang.org/x/oauth2"
)

var (
	user    = flag.String("user", "google", "username target")
	perpage = flag.String("perpage", "100000", "get list repo perpage")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("github")},
	)

	tc := oauth2.NewClient(ctx, ts)
	// service.GetList(tc, *user, *perpage)
	repos, _ := service.GetList(tc, *user, *perpage)
	service.GetListCommits(repos)
}
