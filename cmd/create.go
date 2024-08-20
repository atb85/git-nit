/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates nits in the current git repository given the changed files",
	RunE: func(cmd *cobra.Command, args []string) error {
    return addNits([]string{"cmd/create.go"})
	},
}

// addNits adds the nit comment string to the given files
func addNits(fls []string) error {
  for _, name := range fls {

    f, err := os.Open(name)
    if err != nil {
      return err
    }
    defer f.Close()

    scn := bufio.NewScanner(f)

    cmt := generateComment(name, getNit())

    var mod []string

    for scn.Scan() {
      cur := scn.Text()
      mod = append(mod, cur)
      if shouldAddNit() {
        mod = append(mod, cmt)
      }
    }
    if err := scn.Err(); err != nil {
      return err
    }

    err = os.WriteFile("out.go", []byte(strings.Join(mod, "\n")), 0644)
    if err != nil {
      return err
    }
  }

  return nil
}

// getNit creates a nit (from input)
func getNit() string {
  return "nit"
}

func shouldAddNit() bool {
  ch := rand.Float32()
  return ch < 0.01
}

func init() {
	rootCmd.AddCommand(createCmd)
}
