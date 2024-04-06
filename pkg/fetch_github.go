package pkg

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v61/github"
	"golang.org/x/oauth2"
)

func FetchGithub(account string) string {
	client := github.NewClient(nil)
	ctx := context.Background()

	_, _, err := client.Organizations.List(ctx, account, nil)

	opt := &github.RepositoryListByUserOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByUser(ctx, account, opt)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintln(repos)
}

func oauth_token(token string) *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return tc

}
