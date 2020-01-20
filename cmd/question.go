package cmd

import (

	//"github.com/gidor/pd/dlg"
	"github.com/gen2brain/dlgs"
	// "github.com/martinlindhe/inputbox"
	"github.com/spf13/cobra"
)

func init() {
	// askCmd.Flags().StringVarP(&msgtitle, "title", "t", "", "dialog title")
	// askCmd.Flags().StringVarP(&message, "msg", "m", "", "message to display")
	// askCmd.Flags().StringVarP(&defvalue, "def", "d", "", "valore iniziale")

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
