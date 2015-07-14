package main

import (
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

var (
	org          = flag.String("org", "", "GitHub Organization to list teams for")
	access_token = flag.String("access-token", "", "Access key to use to authenticate with the GitHub API")
)

func main() {

	flag.Parse()

	// exit if access_token is not set
	if *access_token == "" {
		fmt.Println("ERROR: Access token not set! Please supply the -access-token argument to access the GitHub API")
		os.Exit(1)
	}

	// exit if org is not set
	if *org == "" {
		fmt.Println("ERROR: Organization not set! Please supply the -org argument to list teams for your organization.")
		os.Exit(1)
	}

	// set up oauth client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *access_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	// set up github client
	client := github.NewClient(tc)
	options := github.ListOptions{}

	// make api call to get all team_members
	teams, _, _ := client.Organizations.ListTeams(*org, &options)

	fmt.Printf("Getting teams for %s \n", *org)

	// loop through keys and print to stdout
	for _, v := range teams {
		fmt.Printf("Name: %s, ID: %d \n", *v.Name, *v.ID)

	}
}
