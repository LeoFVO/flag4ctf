package cmd

import (
	"fmt"

	flag "github.com/LeoFVO/flag4ctf/pkg"

	"github.com/spf13/cobra"
)

// newCmd represents the build command
var newCmd = &cobra.Command{
	Use:   "make",
	Short: "Make a string flag-compliant",
	Long:  `Provide a string and it build the "flag like" string.`,
	Args: func (cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must provide a string to build")
		}		

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		
		// take the first argument
		initial := args[0]

		builder := flag.NewFlagBuilder()

		builder.SetPrefix(cmd.Flag("prefix").Value.String())
		builder.SetCase(cmd.Flag("case").Value.String())
		builder.SetInitialValue(initial)

		flag := builder.Build()
		fmt.Printf("Here's your flag:\n\n%s",flag)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("prefix", "p", "FLAG", "The prefix to use")
	newCmd.Flags().StringP("case", "c", "preserve", "Case options: preserve, upper, lower")
}
