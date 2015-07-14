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
	get          = flag.String("get", "keys", "-get keys or -get teams")
	org          = flag.String("org", "", "GitHub organization to use when looking up team id's")
	team_id      = flag.Int("team-id", 0, "GitHub Team ID to pull keys from")
)

func get_teams(client github.Client) {

	// exit if org is not set
	if *org == "" {
		fmt.Println("ERROR: Organization not set! Please supply the -org argument to list teams for your organization.")
		os.Exit(1)
	}

	options := github.ListOptions{}

	// make api call to get all team_members
	teams, _, _ := client.Organizations.ListTeams(*org, &options)

	fmt.Printf("Getting teams for %s \n", *org)

	// loop through keys and print to stdout
	for _, v := range teams {
		fmt.Printf("Name: %s, ID: %d \n", *v.Name, *v.ID)

	}
}

func get_keys(client github.Client) {
	// exit if access_token is not set
	if *team_id == 0 {
		fmt.Println("ERROR: Team ID is not set! Please supply the -team-id argument to get your teams keys")
		os.Exit(1)
	}
	// empty set of options
	options := github.OrganizationListTeamMembersOptions{}

	// make api call to get all team_members<%  %>
	team_members, _, _ := client.Organizations.ListTeamMembers(*team_id, &options)

	// loop through keys and print to stdout
	for _, v := range team_members {
		url := fmt.Sprintf("https://github.com/%s.keys", *v.Login)
		resp, _ := http.Get(url)
		contents, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(contents))
	}

}

func get_help() {
	flag.Usage()
}

func main() {

	flag.Parse()

	// exit if access_token is not set
	if *access_token == "" {
		fmt.Println("ERROR: Access token not set! Please supply the -access-token argument to access the GitHub API")
		os.Exit(1)
	}

	// set up oauth client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *access_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	// set up github client
	client := github.NewClient(tc)

	switch {
	case *get == "keys":
		get_keys(*client)
	case *get == "teams":
		get_teams(*client)
	default:
		get_help()
	}

}
