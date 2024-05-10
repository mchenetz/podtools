package utils

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func GetValue(cmd cobra.Command, path string) string {
	if cmd.Flags().Changed(path) {
		return cmd.Flag(path).Value.String()
	} else {
		return viper.GetString(path)
	}
}

func SplitFilePath(path string) (string, string) {
	dir, filename := filepath.Split(path)
	return dir, filename
}
