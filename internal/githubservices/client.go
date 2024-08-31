package githubservices

import (
	"context"

	"github.com/google/go-github/v63/github"
	"golang.org/x/oauth2"
)

type services struct {
	pr   PullRequestServicer
	git  GitServicer
	repo RepositoriesServicer
}

type PullRequestServicer interface {
	ListFiles(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.CommitFile, *github.Response, error)
	ListReviewComments(ctx context.Context, owner string, repo string, number int, reviewID int64, opts *github.ListOptions) ([]*github.PullRequestComment, *github.Response, error)
	ListReviews(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.PullRequestReview, *github.Response, error)
}

type GitServicer interface {
	GetRef(ctx context.Context, owner string, repo string, ref string) (*github.Reference, *github.Response, error)
	GetCommit(ctx context.Context, owner string, repo string, sha string) (*github.Commit, *github.Response, error)
	GetTree(ctx context.Context, owner string, repo string, sha string, recursive bool) (*github.Tree, *github.Response, error)
	CreateTree(ctx context.Context, owner string, repo string, baseTree string, entries []*github.TreeEntry) (*github.Tree, *github.Response, error)
	CreateCommit(ctx context.Context, owner string, repo string, commit *github.Commit, opts *github.CreateCommitOptions) (*github.Commit, *github.Response, error)
	UpdateRef(ctx context.Context, owner string, repo string, ref *github.Reference, force bool) (*github.Reference, *github.Response, error)
}

type RepositoriesServicer interface {
	GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error)
}

type PullRequestServices struct{}

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
			Body:     github.String("nit"),
		},
		&github.PullRequestComment{
			DiffHunk: github.String("normal diff\n\nright here"),
			Body:     github.String("nit"),
		},
		&github.PullRequestComment{
			DiffHunk: github.String("normal diff\n\nright here"),
			Body:     github.String("normal comment"),
		},
		&github.PullRequestComment{
			DiffHunk: github.String("another nit\n\nright here"),
			Body:     github.String("normal comment"),
		},
	}

	return cmts, &github.Response{}, nil
}

func (s *PullRequestServices) ListReviews(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.PullRequestReview, *github.Response, error) {
	rvws := []*github.PullRequestReview{
		&github.PullRequestReview{
			State: github.String("APPROVED"),
			ID:    github.Int64(9),
		},
		&github.PullRequestReview{
			State: github.String(""),
			ID:    github.Int64(7),
		},
	}

	return rvws, &github.Response{}, nil
}

func NewClient(tkn string) *services {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: tkn})

	tc := oauth2.NewClient(ctx, ts)
	clnt := github.NewClient(tc)

	return &services{
		pr: clnt.PullRequests,
	}
}

func newFakeClient() *services {

	return &services{
		pr: &PullRequestServices{},
	}

}
