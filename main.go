package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

func main() {
	fmt.Println("Hello")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(repos)
}
