package githubservices

import (
	"context"
	"errors"
	"git-nit/internal/nits"
	"slices"
	"strings"

	"github.com/google/go-github/v63/github"
)

type Pr struct {
	Owner  string
	Repo   string
	Branch string
	Number int
	Ctx    context.Context
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
	if err != nil {
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

func (p *Pr) addNits(clnt *services, files []string) error {

	// Get the reference for the branch
	ref, _, err := clnt.git.GetRef(p.Ctx, p.Owner, p.Repo, "refs/heads/"+p.Branch)
	if err != nil {
		return err
	}

	// Get the tree for the latest commit
	commit, _, err := clnt.git.GetCommit(p.Ctx, p.Owner, p.Repo, ref.GetObject().GetSHA())
	if err != nil {
		return err
	}

	// Get the tree
	tree, _, err := clnt.git.GetTree(p.Ctx, p.Owner, p.Repo, commit.GetTree().GetSHA(), true)
	if err != nil {
		return err
	}

	// Create a new tree with modified content
	var newTree []*github.TreeEntry
	for _, entry := range tree.Entries {
		if entry.GetType() == "blob" && slices.Contains(files, entry.GetPath()) {

			content, _, _, err := clnt.repo.GetContents(p.Ctx, p.Owner, p.Repo, entry.GetPath(), &github.RepositoryContentGetOptions{Ref: p.Branch})
			if err != nil {
				return err
			}

			file, err := content.GetContent()
			if err != nil {
				return err
			}

			// Add comment to the content
			nit := nits.NewNit(entry.GetPath())
			newContent, err := nits.AddNit([]byte(file), nit, 0.03)
			if err != nil {
				return err
			}

			newEntry := &github.TreeEntry{
				Path:    github.String(entry.GetPath()),
				Mode:    github.String("100644"),
				Type:    github.String("blob"),
				Content: github.String(string(newContent)),
			}
			newTree = append(newTree, newEntry)
		} else {
			newTree = append(newTree, entry)
		}
	}

	// Create a new tree
	newTreeObject, _, err := clnt.git.CreateTree(p.Ctx, p.Owner, p.Repo, tree.GetSHA(), newTree)
	if err != nil {
		return err
	}

	cmt := &github.Commit{
		Message: github.String("Add TODO comment"),
		Tree:    newTreeObject,
		Parents: []*github.Commit{{SHA: commit.SHA}},
	}

	// Create a new commit
	newCommit, _, err := clnt.git.CreateCommit(p.Ctx, p.Owner, p.Repo, cmt, nil)
	if err != nil {
		return err
	}

	newRef := &github.Reference{
		Ref: github.String("refs/heads/" + p.Branch),
		Object: &github.GitObject{
			SHA: newCommit.SHA,
		},
	}

	// Update the reference to point to the new commit
	_, _, err = clnt.git.UpdateRef(p.Ctx, p.Owner, p.Repo, newRef, false)
	if err != nil {
		return err
	}

	return nil
}
