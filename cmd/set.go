/*
Copyright © 2021 Gianni Doria  <gianni.doria@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	setCmd.Flags().StringVarP(&defvalue, "def", "d", "", "valore iniziale")

	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set <value>",
	Short: "set value in config  file",
	Long: `set value in config. 
	Configuration may be in json or yaml format`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(parName) > 0 {
			iniCfg()
			defer finalizeCfg()
			setParam(parName, args[0])
		}
	},
}
