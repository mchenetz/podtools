/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Save Preset in Config File",
	Long:  `Save Preset in Config File`,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flags().Changed("read") {
			for _, i := range viper.AllKeys() {
				fmt.Println(i+":", viper.Get(i))
			}
		} else {
			viper.WriteConfig()
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("read", "r", false, "Read Config File")
	configCmd.PersistentFlags().StringP("inputfile", "i", "./input.mp4", "Convert Input file")
	configCmd.PersistentFlags().StringP("outputfile", "o", "./output.mp3", "Convert Output file")
	configCmd.PersistentFlags().StringP("transcribefile", "t", "./transcribe.mp3", "Transcribe File")
	configCmd.PersistentFlags().StringP("defaultfolder", "f", "", "Default Folder")
	configCmd.PersistentFlags().StringP("apikey", "a", viper.GetString("apikey"), "OpenAI API Key")
	viper.BindPFlag("inputfile", configCmd.PersistentFlags().Lookup("inputfile"))
	viper.BindPFlag("outputfile", configCmd.PersistentFlags().Lookup("outputfile"))
	viper.BindPFlag("defaultfolder", configCmd.PersistentFlags().Lookup("defaultfolder"))
	viper.BindPFlag("apikey", configCmd.PersistentFlags().Lookup("apikey"))
	viper.BindPFlag("transcribefile", configCmd.PersistentFlags().Lookup("transcribefile"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
