package main

import (
	"context"
	service "findkey/services"
	"flag"
	"fmt"
	"os"

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

	repos := service.GetList(tc, *user)
	fmt.Println(repos)
}
