package githubservices

import (
	"context"
	"fmt"

	"github.com/google/go-github/v63/github"
)

func addCommentsToBranch(ctx context.Context, clnt *services, own, repo, branch string) error {
	// Get the reference for the branch
	ref, _, err := clnt.git.GetRef(ctx, own, repo, "refs/heads/"+branch)
	if err != nil {
		return fmt.Errorf("error getting ref: %v", err)
	}

	// Get the tree for the latest commit
	commit, _, err := clnt.git.GetCommit(ctx, own, repo, ref.GetObject().GetSHA())
	if err != nil {
		return fmt.Errorf("error getting commit: %v", err)
	}

	// Get the tree
	tree, _, err := clnt.git.GetTree(ctx, own, repo, commit.GetTree().GetSHA(), true)
	if err != nil {
		return fmt.Errorf("error getting tree: %v", err)
	}

	// Create a new tree with modified content
	var newTree []*github.TreeEntry
	for _, entry := range tree.Entries {
		if entry.GetType() == "blob" && entry.GetPath() == "main.go" { // Example: modifying main.go
			content, _, _, err := clnt.repo.GetContents(ctx, own, repo, entry.GetPath(), &github.RepositoryContentGetOptions{Ref: branch})
			if err != nil {
				return fmt.Errorf("error getting file content: %v", err)
			}

			file, err := content.GetContent()
			if err != nil {
				return fmt.Errorf("error decoding content: %v", err)
			}

			// Add comment to the content
			newContent := []byte("// TODO: Review this file for potential improvements\n\n" + string(file))

			newEntry := &github.TreeEntry{
				Path:    github.String(*entry.Path),
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
	newTreeObject, _, err := clnt.git.CreateTree(ctx, own, repo, tree.GetSHA(), newTree)
	if err != nil {
		return fmt.Errorf("error creating new tree: %v", err)
	}

	cmt := &github.Commit{
		Message: github.String("Add TODO comment"),
		Tree:    newTreeObject,
		Parents: []*github.Commit{{SHA: commit.SHA}},
	}

	// Create a new commit
	newCommit, _, err := clnt.git.CreateCommit(ctx, own, repo, cmt, nil)
	if err != nil {
		return fmt.Errorf("error creating new commit: %v", err)
	}

	// Update the reference to point to the new commit
	_, _, err = clnt.git.UpdateRef(ctx, own, repo, &github.Reference{
		Ref: github.String("refs/heads/" + branch),
		Object: &github.GitObject{
			SHA: newCommit.SHA,
		},
	}, false)
	if err != nil {
		return fmt.Errorf("error updating ref: %v", err)
	}

	return nil
}
