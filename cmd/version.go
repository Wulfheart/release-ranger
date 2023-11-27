package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	version   string = "dev"
	goversion string = "n/a"
	commit    string = "n/a"
	built     string = "n/a"
	build     string = "n/a"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the current version",
	Run: func(cmd *cobra.Command, args []string) {
		data := [][]string{
			{"Version:", version},
			{"Go version:", goversion},
			{"Git commit:", commit},
			{"Built", built},
			{"Build", build},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.SetBorder(false)
		table.SetTablePadding("\t") // pad with tabs
		table.SetNoWhiteSpace(true)
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
