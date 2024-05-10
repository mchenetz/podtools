/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"podtools/ffmpeg"
	"podtools/openai"
	"podtools/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// transcribeCmd represents the transcribe command
var transcribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "Transcribes your audio file using OpenAI.",
	Long:  `Transcribes your audio file using OpenAI.`,
	Run: func(cmd *cobra.Command, args []string) {
		var tfile string
		var dir string
		tfile = utils.GetValue(*cmd, "tinputfile")
		dir = utils.GetValue(*cmd, "defaultfolder")
		fmt.Println("path:", dir+tfile)
		fmt.Println("transcribe called")
		apikey := viper.GetString("apikey")
		ffmpeg.ConvertToMp3(dir+tfile, dir+"transcribe.mp3", "32k")
		client := openai.Connect(apikey)
		openai.Transcribe(client, dir+tfile)
		os.Remove(dir + "transcribe.mp3")
	},
}

func init() {
	transcribeCmd.Flags().StringP("tinputfile", "i", "./transcribe.mp3", "File path")
	transcribeCmd.Flags().StringP("defaultfolder", "f", "", "Default Folder")
	rootCmd.AddCommand(transcribeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transcribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	transcribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
