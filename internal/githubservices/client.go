package githubservices

import (
	"context"

	"github.com/google/go-github/v63/github"
	"golang.org/x/oauth2"
)

type services struct {
  pr PullRequestServices
}

type PullRequestServices struct {}

func (s *PullRequestServices) ListFiles(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.CommitFile, *github.Response, error) {
  return []*github.CommitFile{
    &github.CommitFile{
      Filename: github.String("test.go"),
    },
    &github.CommitFile{
      Filename: github.String("work.go"),
    },
  }, &github.Response{}, nil
}

func (s *PullRequestServices) ListReviewComments(ctx context.Context, owner string, repo string, number int, reviewID int64, opts *github.ListOptions) ([]*github.PullRequestComment, *github.Response, error) {

  cmts := []*github.PullRequestComment{
    &github.PullRequestComment{
      DiffHunk: github.String("this is a nit\n\nright here"),
      Body: github.String("nit"),
    },
    &github.PullRequestComment{
      DiffHunk: github.String("normal diff\n\nright here"),
      Body: github.String("nit"),
    },
    &github.PullRequestComment{
      DiffHunk: github.String("normal diff\n\nright here"),
      Body: github.String("normal comment"),
    },
    &github.PullRequestComment{
      DiffHunk: github.String("another nit\n\nright here"),
      Body: github.String("normal comment"),
    },
  }

  return cmts, &github.Response{}, nil
}

func (s *PullRequestServices) ListReviews(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.PullRequestReview, *github.Response, error) {
  rvws := []*github.PullRequestReview{
    &github.PullRequestReview{
      State: github.String("APPROVED"),
      ID: github.Int64(9),
    },
    &github.PullRequestReview{
      State: github.String(""),
      ID: github.Int64(7),
    },
  }

  return rvws, &github.Response{}, nil
}

func newClient(tkn string) *services {

  ctx := context.Background()
  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: tkn})

  tc := oauth2.NewClient(ctx, ts)
  clnt := github.NewClient(tc)

  return &services{
    pr: clnt.PullRequests,
  }
}

func newFakeClient() *services {

}
