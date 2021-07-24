/*
Copyright © 2021 Gianni Doria  <gianni.doria@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// var title string

func init() {
	rootCmd.AddCommand(dirCmd)
}

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "manage dir parameter ",
	Long:  `manage dir parameter `,
	Run: func(cmd *cobra.Command, args []string) {
		iniCfg()
		defer finalizeCfg()
		fname, err := dlgDir(title)
		check(err)
		setPathParam(parName, fname)
	},
}
