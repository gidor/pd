/*
Copyright © 2021 Gianni Doria  <gianni.doria@gmail.com>
*/
package cmd

import (
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
		if len(parName) > 0 {
			iniCfg()
			defer finalizeCfg()
			ok := dlgYesNo(message, title) //dlgs.Question(title, message, true)
			if ok {
				setParam(parName, "ok")
			}
		}
	},
}
