/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// majorCmd represents the major command
var majorCmd = &cobra.Command{
	Use:   "major",
	Short: "Increment the major version and push",
	Run: func(cmd *cobra.Command, args []string) {
		versions, err := releaser.Retrieve()
		if err != nil {
			panic(err)
		}
		latestVersion := versions[0]
		newVersion := latestVersion.IncrementMajor()
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
	rootCmd.AddCommand(majorCmd)
}
