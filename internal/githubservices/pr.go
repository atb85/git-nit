package githubservices

import (
	"context"
	"errors"

	"github.com/google/go-github/v63/github"
)

type pr struct{
  owner  string
  repo   string
  number int
  ctx context.Context
}

func (p *pr) getFiles(client *github.Client) ([]string, error) {
  fs, _, err := client.PullRequests.ListFiles(p.ctx, p.owner, p.repo, p.number, nil)
  if err != nil {
    return []string{}, err
  }

  var names []string

  for _, f := range fs {
    n := f.GetFilename()
    if n != "" {
      names = append(names, f.GetFilename())
    }
  }

  return names, nil
}

func (p *pr) getApprovedReviews(client *github.Client) ([]*github.PullRequestReview, error) {
  rvws, _, err := client.PullRequests.ListReviews(p.ctx, p.owner, p.repo, p.number, nil)
  if err != nil{
    return []*github.PullRequestReview{}, err
  }

  var approvals []*github.PullRequestReview
  for _, rvw := range rvws {
    if rvw.GetState() == "APPROVED" {
      approvals = append(approvals, rvw)
    }
  }
  if approvals == nil {
    return []*github.PullRequestReview{}, errors.New("no approved reviews")
  }

  return approvals, nil
}

func (p *pr) getNitComments(client *github.Client, rvw *github.PullRequestReview) {
  cmts, _, err := client.PullRequests.ListReviewComments(p.ctx, p.owner, p.repo, p.number, rvw.GetID())
  for _, cmt := range cmts {
    cmt.get
  }
}
