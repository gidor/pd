/*
Copyright © 2021 Gianni Doria  <gianni.doria@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&configName, "cfg", "c", "cfg", "configuration name ")
	rootCmd.PersistentFlags().StringVarP(&parName, "name", "n", "", "parameter name")
	rootCmd.PersistentFlags().Int8VarP(&convDelimiter, "win", "w", win, "convert path delimiter using win (backslash)")
	rootCmd.PersistentFlags().Int8VarP(&convDelimiter, "unix", "u", unix, "convert path delimiter using unix (slash)")
	rootCmd.PersistentFlags().Int8VarP(&convDelimiter, "path", "p", unix, "convert path delimiter using unix (slash)")
	rootCmd.PersistentFlags().BoolVarP(&jyout, "jy", "2", false, "dual config output")
	rootCmd.PersistentFlags().StringVarP(&parValue, "value", "v", "", "parameter value")

	rootCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "dialog title")
	rootCmd.PersistentFlags().StringVarP(&message, "msg", "m", "", "message to display")

	basePath = rootPath()
}

var rootCmd = &cobra.Command{
	Use:   "pd",
	Short: "pd path and dialog",
	Long:  `pd dialog for simple interactive procedure`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(parName) > 0 {
			iniCfg()
			defer finalizeCfg()
			setParam(parName, parValue)
		}
	},
}

// Execute  execute comamdd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
