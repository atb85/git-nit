/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-nit",
	Short: "the code review assurance tool",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cmd.PersistentFlags().Int("number", -1, "the pull request number to validate reviews on")
	cmd.PersistentFlags().String("owner", "", "repo owner - organisation or username")
	cmd.PersistentFlags().String("repo", "", "the name of the repository")

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

}
