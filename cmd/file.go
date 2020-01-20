package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/sqweek/dialog"
)

var filter string

func init() {
	fileCmd.Flags().StringVarP(&filter, "filter", "f", "", "filter in the form name|ext")
	// fileCmd.Flags().StringVarP(&filetitle, "title", "t", "", "dialog title")

	rootCmd.AddCommand(fileCmd)
}

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "manage file parameter ",
	Long: `manage file parameter
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(parName) > 0 {
			fdialog := dialog.File().Title(title).Filter("All files", "*")
			sfilter := strings.Split(filter, "|")
			if len(sfilter) >= 2 {
				fdialog.Filter(sfilter[0], sfilter[1])
			}
			f, err := fdialog.Load()
			check(err)
			setParam(parName, f)
		}
	},
}
