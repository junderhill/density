package cmd

import (
	"fmt"
	"os"

	"github.com/junderhill/density/internal/cmd/shoot"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "density",
	Short: "Density provides a score based indication of landscape photography conditions at a given location",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	rootCmd.AddCommand(shoot.ShootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
