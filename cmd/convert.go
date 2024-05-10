/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"podtools/ffmpeg"
	"podtools/utils"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts a video to mp3.",
	Long:  `Converts a video to mp3.`,
	Run: func(cmd *cobra.Command, args []string) {
		var dir string
		var in string
		var out string
		in = utils.GetValue(*cmd, "inputfile")
		out = utils.GetValue(*cmd, "outputfile")
		dir = utils.GetValue(*cmd, "defaultfolder")
		if utils.FileExists(dir + out) {
			println("File exists")
		} else {
			ffmpeg.ConvertToMp3(dir+in, dir+out, "192k")
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringP("inputfile", "i", "./input.mp4", "Convert Input file")
	convertCmd.Flags().StringP("outputfile", "o", "./output.mp3", "Convert Output file")
	convertCmd.Flags().StringP("defaultfolder", "f", "", "Default Folder")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
