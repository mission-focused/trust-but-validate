/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type flags struct {
	Source string // -s --source
	Tag    string // -t --tag
}

var opts = &flags{}

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a go checksum",
	Long:  `Perform validation of a go module checksum against the go sum db and actual hashing of the remote repository`,
	Run: func(cmd *cobra.Command, args []string) {
		if opts.Source == "" {
			fmt.Println("source is required")
			os.Exit(1)
		}

		if opts.Tag == "" {
			fmt.Println("tag is required")
			os.Exit(1)
		}

		fmt.Println("validate called")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	validateCmd.Flags().StringVarP(&opts.Source, "source", "s", "", "Source of the module")
	validateCmd.Flags().StringVarP(&opts.Tag, "tag", "t", "", "Tag of the module")

}
