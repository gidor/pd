/*
Copyright © 2021 Gianni Doria  <gianni.doria@gmail.com>
*/
package cmd

import (
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
		if len(parName) > 0 {
			iniCfg()
			defer finalizeCfg()
			value, err := promptMessage(message)
			// value, ok, err := dlgs.Entry(title, message, defvalue)
			check(err)
			setParam(parName, value)
		}
	},
}
