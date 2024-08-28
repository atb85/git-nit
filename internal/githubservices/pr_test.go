package githubservices

import (
	"context"
	"testing"

	"github.com/google/go-github/v63/github"
	"github.com/stretchr/testify/assert"
)

func Test_GetValidNitPicks(t *testing.T) {

  clnt := newFakeClient()

  p := &Pr{
    Owner: "test",
    Repo: "test",
    Number: 0,
    Ctx: context.Background(),
  }
  
  rvw := &github.PullRequestReview{
    ID: github.Int64(744),
  }

  nits, err := p.GetValidNitPicks(clnt, rvw)
  assert.NoError(t, err)
  assert.Equal(t, 1, nits)
}

func Test_GetApprovedReviews(t *testing.T) {

  clnt := newFakeClient()

  p := &Pr{
    Owner: "test",
    Repo: "test",
    Number: 0,
    Ctx: context.Background(),
  }

  reviews, err := p.GetApprovedReviews(clnt)
  assert.NoError(t, err)
  for _, rvw := range reviews {
    assert.Equal(t, "APPROVED", rvw.GetState())
  }
}

func Test_GetFiles(t *testing.T) {

  clnt := newFakeClient()

  p := &Pr{
    Owner: "test",
    Repo: "test",
    Number: 0,
    Ctx: context.Background(),
  }

  fs, err := p.GetFiles(clnt)
  assert.NoError(t, err)
  for _, f := range fs {
    assert.NotEmpty(t, f)
  }
}
