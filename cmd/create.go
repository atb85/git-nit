package cmd

import (
	"bufio"
	"fmt"
	"git-nit/internal/nits"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func shouldAddNit() bool {
	ch := rand.Float32()
	return ch < 0.01
}

func getNit() string {
	return "nit"
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates nits in the current git repository given the changed files",
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

		fmt.Println(num, own, rpo)

		return addNits("")
	},
}

// addNits adds the nit comment string to the given files
func addNits(fls ...string) error {

	for _, name := range fls {
		f, err := os.Open(name)
		if err != nil {
			return err
		}
		defer f.Close()

		scn := bufio.NewScanner(f)

		cmt := nits.GenerateComment(name, getNit())

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

		err = os.WriteFile(name, []byte(strings.Join(mod, "\n")), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)
}
