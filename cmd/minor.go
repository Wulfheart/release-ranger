/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gookit/color"

	"github.com/spf13/cobra"
)

// minorCmd represents the minor command
var minorCmd = &cobra.Command{
	Use: "minor",
	Run: func(cmd *cobra.Command, args []string) {
		versions, err := releaser.Retrieve()
		if err != nil {
			panic(err)
		}
		latestVersion := versions[0]
		newVersion := latestVersion.IncrementMinor()
		fmt.Println("Incrementing " + latestVersion.String() + " to " + newVersion.String())

		if !dryRun {
			err = releaser.Create(newVersion)
			if err != nil {
				panic(err)
			}
		}

		color.Greenln("Created " + newVersion.String())
	},
}

func init() {
	rootCmd.AddCommand(minorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// minorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// minorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
