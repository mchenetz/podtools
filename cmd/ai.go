/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"podtools/openai"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// aiCmd represents the ai command
var aiCmd = &cobra.Command{
	Use:   "ai",
	Short: "Submit a query to AI",
	Long:  `Submit a query to AI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ai called")
		client := openai.Connect(viper.GetString("apikey"))
		fn, err := cmd.Flags().GetString("aifilename")
		if err != nil {
			fmt.Println(err)
		}
		query, err := cmd.Flags().GetString("query")
		if err != nil {
			fmt.Println(err)
		}
		openai.SubmitToGPT4(client, fn, query)
	},
}

func init() {
	rootCmd.AddCommand(aiCmd)
	aiCmd.Flags().StringP("aifilename", "i", "", "File to submit")
	aiCmd.Flags().StringP("query", "q", "", "Query to submit")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
