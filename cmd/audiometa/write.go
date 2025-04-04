package main

import (
	"fmt"

	"github.com/gcottom/audiometa"
	"github.com/spf13/cobra"
)

var (
	artist      string
	albumArtist string
	album       string
	comment     string
	composer    string
	genre       string
	title       string
	track       string
	year        string
	date        string
	bpm         string
	cover       string
)

func init() {
	flags := writeCmd.Flags()
	flags.StringVar(&artist, "artist", "", "Artist name")
	flags.StringVar(&albumArtist, "albumartist", "", "Album artist name")
	flags.StringVar(&album, "album", "", "Album name")
	flags.StringVar(&comment, "comment", "", "Comment")
	flags.StringVar(&composer, "composer", "", "Composer")
	flags.StringVar(&genre, "genre", "", "Genre")
	flags.StringVar(&title, "title", "", "Title")
	flags.StringVar(&track, "track", "", "Track number")
	flags.StringVar(&year, "year", "", "Year")
	flags.StringVar(&date, "date", "", "Date")
	flags.StringVar(&bpm, "bpm", "", "BPM")
	flags.StringVar(&cover, "cover", "", "Cover art filename")
	rootCmd.AddCommand(writeCmd)
}

func writeTags(file string) error {
	tag, err := audiometa.OpenTag(file)
	if err != nil {
		return err
	}
	if artist != "" {
		tag.SetArtist(artist)
	}
	if albumArtist != "" {
		tag.SetAlbumArtist(albumArtist)
	}
	if album != "" {
		tag.SetAlbum(album)
	}
	if comment != "" {
		tag.SetComments(comment)
	}
	if composer != "" {
		tag.SetComposer(composer)
	}
	if genre != "" {
		tag.SetGenre(genre)
	}
	if title != "" {
		tag.SetTitle(title)
	}
	if track != "" {
		tag.SetTrack(track)
	}
	if year != "" {
		tag.SetYear(year)
	}
	if date != "" {
		tag.SetDate(date)
	}
	if bpm != "" {
		tag.SetBPM(bpm)
	}
	if cover != "" {
		tag.SetAlbumArtFromFilePath(cover)
	}
	// TODO: add support for additional ogg tags
	// tag.SetAdditionalTag(strings.ToUpper(cmdTag), writeTag)
	err = tag.Save()
	return err
}

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write some tags",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			err := writeTags(file)
			if err != nil {
				fmt.Printf("%s: Tags saved failure:\n%s\n", file, err)
			} else {
				fmt.Printf("Tags saved successfully: %s\n", file)
			}

		}
	},
}
