/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"git-nit/internal/githubservices"
	"os"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var cmd = &cobra.Command{
	Use:   "validate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
    num, err := cmd.Flags().GetInt("number")
    if err != nil {
      return err
    }

    own, err := cmd.Flags().GetString("owner")
    if err != nil {
      return err
    }

    rpo, err := cmd.Flags().GetString("repo")
    if err != nil {
      return err
    }

    pull := &githubservices.Pr{
      Owner: own,
      Repo: rpo,
      Number: num,
      Ctx: context.Background(),
    }

    tkn := os.Getenv("GITHUB_TOKEN")
    if tkn == "" {
      return errors.New("no github token set - set environment variable GITHUB_TOKEN")
    }

    clnt := githubservices.NewClient(tkn)

    apps, err := pull.GetApprovedReviews(clnt)
    if err != nil {
      return err
    }

    for _, rvw := range apps {
      nits, err := pull.GetValidNitPicks(clnt, rvw)
      if err != nil {
        return err
      }
    }

    return nil
	},
}

func init() {
	rootCmd.AddCommand(cmd)

  cmd.Flags().Int("number", -1, "the pull request number to validate reviews on")
  cmd.Flags().String("owner", "", "repo owner - organisation or username")
  cmd.Flags().String("repo", "", "the name of the repository")

  req := []string{
    "number",
    "owner",
    "repo",
  }

  for _, flag := range req {
    err := cmd.MarkFlagRequired(flag)
    if err != nil {
      panic(err)
    }
  }


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
