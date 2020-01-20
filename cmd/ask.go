package cmd

import (

	//"github.com/gidor/pd/dlg"
	"github.com/gen2brain/dlgs"
	// "github.com/martinlindhe/inputbox"
	"github.com/spf13/cobra"
)

var defvalue string

func init() {
	// askCmd.Flags().StringVarP(&msgtitle, "title", "t", "", "dialog title")
	// askCmd.Flags().StringVarP(&message, "msg", "m", "", "message to display")
	askCmd.Flags().StringVarP(&defvalue, "def", "d", "", "valore iniziale")

	rootCmd.AddCommand(askCmd)
}

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "ask question to user",
	Long:  `ask question to user and store answare as a key value`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(parName) > 0 {
			value, ok, err := dlgs.Entry(title, message, defvalue)
			if ok {
				check(err)
				setParam(parName, value)
			}
			// fmt.Println(ok)
			// fmt.Println(value)
		}
	},
}
