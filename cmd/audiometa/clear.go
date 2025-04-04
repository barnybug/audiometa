package main

import (
	"fmt"

	"github.com/gcottom/audiometa"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clearCmd)
}

func clearTags(file string) error {
	tag, err := audiometa.OpenTag(file)
	if err != nil {
		return err
	}
	tag.ClearAllTags(false)
	err = tag.Save()
	return err
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all tags",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			err := clearTags(file)
			if err != nil {
				fmt.Printf("%s: Tags clear failure:\n%s\n", file, err)
			} else {
				fmt.Printf("Tags cleared successfully: %s\n", file)
			}

		}
	},
}
