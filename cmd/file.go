/*
Copyright © 2021 Gianni Doria  <gianni.doria@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var filter string

func init() {
	fileCmd.Flags().StringVarP(&filter, "filter", "f", "", "filter in the form name|ext")

	rootCmd.AddCommand(fileCmd)
}

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "manage file parameter ",
	Long:  `manage file parameter`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(parName) > 0 {
			iniCfg()
			defer finalizeCfg()
			fname, err := dlgFile(title)
			check(err)
			setPathParam(parName, fname)
		}
	},
}
