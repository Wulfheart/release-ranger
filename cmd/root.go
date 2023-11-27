/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gookit/color"
	"github.com/wulfheart/release-ranger/core"
	"os"

	"github.com/spf13/cobra"
)

var releaser core.Releaser
var dryRun bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rer",
	Short: "",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "Dry run")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if dryRun {
			color.BgRed.Println("Dry run enabled")
		}
	}
	releaser = core.GitReleaser{}
}
