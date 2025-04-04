package main

import (
	"fmt"

	"github.com/gcottom/audiometa"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

func readTags(file string) error {
	tag, err := audiometa.OpenTag(file)
	if err != nil {
		return err
	}
	printedTags := 0
	fmt.Printf("File: %s\n", file)
	if tag.Artist() != "" {
		fmt.Printf("Artist: %s\n", tag.Artist())
		printedTags++
	}
	if tag.AlbumArtist() != "" {
		fmt.Printf("AlbumArtist: %s\n", tag.AlbumArtist())
		printedTags++
	}
	if tag.Album() != "" {
		fmt.Printf("Album: %s\n", tag.Album())
		printedTags++
	}
	if tag.BPM() != "" {
		fmt.Printf("BPM: %s\n", tag.BPM())
		printedTags++
	}
	if tag.Comments() != "" {
		fmt.Printf("Comment: %s\n", tag.Comments())
		printedTags++
	}
	if tag.Composer() != "" {
		fmt.Printf("Composer: %s\n", tag.Composer())
		printedTags++
	}
	if tag.CopyrightMsg() != "" {
		fmt.Printf("Copyright: %s\n", tag.CopyrightMsg())
		printedTags++
	}
	if tag.Date() != "" {
		fmt.Printf("Date: %s\n", tag.Date())
		printedTags++
	}
	if tag.EncodedBy() != "" {
		fmt.Printf("EncodedBy: %s\n", tag.EncodedBy())
		printedTags++
	}
	if tag.Genre() != "" {
		fmt.Printf("Genre: %s\n", tag.Genre())
		printedTags++
	}
	if tag.Language() != "" {
		fmt.Printf("Language: %s\n", tag.Language())
		printedTags++
	}
	if tag.Length() != "" {
		fmt.Printf("Length: %s\n", tag.Length())
		printedTags++
	}
	if tag.Lyricist() != "" {
		fmt.Printf("Lyricist: %s\n", tag.Lyricist())
		printedTags++
	}
	if tag.PartOfSet() != "" {
		fmt.Printf("PartOfSet: %s\n", tag.PartOfSet())
		printedTags++
	}
	if tag.Publisher() != "" {
		fmt.Printf("Publisher: %s\n", tag.Publisher())
		printedTags++
	}
	if tag.Track() != "" {
		fmt.Printf("Track: %s\n", tag.Track())
		printedTags++
	}
	if tag.Title() != "" {
		fmt.Printf("Title: %s\n", tag.Title())
		printedTags++
	}
	if tag.Year() != "" {
		fmt.Printf("Year: %s\n", tag.Year())
		printedTags++
	}
	if len(tag.AdditionalTags()) > 0 {
		fmt.Print("\nUnmapped Tags:\n")
		for key, value := range tag.AdditionalTags() {
			fmt.Printf("%s: %s\n", key, value)
			printedTags++
		}
	}
	if printedTags == 0 {
		fmt.Println("No tags found for this file!")
	}
	fmt.Println("--------------------------------")

	return nil
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read tags",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			err := readTags(file)
			if err != nil {
				fmt.Printf("%s: Tags read failure:\n%s\n", file, err)
			}
		}
	},
}
