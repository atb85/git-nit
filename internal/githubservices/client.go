package githubservices

import (
	"context"

	"github.com/google/go-github/v63/github"
	"golang.org/x/oauth2"
)

func newClient(tkn string) *github.Client {

  ctx := context.Background()
  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: tkn})

  tc := oauth2.NewClient(ctx, ts)
  clnt := github.NewClient(tc)

  return clnt
}
