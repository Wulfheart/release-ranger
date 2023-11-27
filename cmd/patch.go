/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gookit/color"

	"github.com/spf13/cobra"
)

// patchCmd represents the patch command
var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Increment the patch version and push",
	Run: func(cmd *cobra.Command, args []string) {
		versions, err := releaser.Retrieve()
		if err != nil {
			panic(err)
		}
		latestVersion := versions[0]
		newVersion := latestVersion.IncrementPatch()
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
	rootCmd.AddCommand(patchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
