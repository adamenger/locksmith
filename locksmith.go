package main

import (
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	access_token = flag.String("access-token", "", "Access key to use to authenticate with the GitHub API")
	team_id      = flag.Int("team-id", 0, "GitHub Team ID to pull keys from")
)

func main() {

	flag.Parse()

	// exit if access_token is not set
	if *access_token == "" {
		fmt.Println("Access token not set! Please supply the -access-token argument to access the GitHub API")
		os.Exit(1)
	}

	// set up oauth client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *access_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	// set up github client
	client := github.NewClient(tc)
	options := github.OrganizationListTeamMembersOptions{}

	// make api call to get all team_members
	team_members, _, _ := client.Organizations.ListTeamMembers(*team_id, &options)

	// loop through keys and print to stdout
	for _, v := range team_members {
		url := fmt.Sprintf("https://github.com/%s.keys", *v.Login)
		resp, _ := http.Get(url)
		contents, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(contents))
	}
}
