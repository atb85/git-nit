package githubservices

import (
	"context"
	"errors"
	"strings"

	"github.com/google/go-github/v63/github"
)

type Pr struct{
  Owner  string
  Repo   string
  Number int
  Ctx context.Context
}

func (p *Pr) GetFiles(client *services) ([]string, error) {
  fs, _, err := client.pr.ListFiles(p.Ctx, p.Owner, p.Repo, p.Number, nil)
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

func (p *Pr) GetApprovedReviews(client *services) ([]*github.PullRequestReview, error) {
  rvws, _, err := client.pr.ListReviews(p.Ctx, p.Owner, p.Repo, p.Number, nil)
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

// GetValidNitPicks returns the number of successful nits that were found in the given review
func (p *Pr) GetValidNitPicks(client *services, rvw *github.PullRequestReview) (int, error) {
  cmts, _, err := client.pr.ListReviewComments(p.Ctx, p.Owner, p.Repo, p.Number, rvw.GetID(), nil)
  if err != nil {
    return 0, err
  }

  sum := 0
  for _, cmt := range cmts {
    txt := strings.ToLower(cmt.GetBody())
    if strings.Contains(txt, "nit") {
      // nit is a placeholder
      nit := "nit"
      if strings.Contains(cmt.GetDiffHunk(), nit) {
        sum++
      }
    }
  }

  return sum, nil
}
