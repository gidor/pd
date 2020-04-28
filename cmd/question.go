package cmd

import (
	"github.com/gen2brain/dlgs"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(questionCmd)
}

var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "ask a Yes/No question  to user",
	Long:  `ask Yes/Noquestion to user and store answare as a key value`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(parName) > 0 {
			ok, err := dlgs.Question(title, message, true)
			check(err)
			if ok {
				setParam(parName, "ok")
			}
		}
	},
}
