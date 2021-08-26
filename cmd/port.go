package cmd

import (
	"github.com/secriy/golire/module"
	"github.com/spf13/cobra"
)

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:   "port [Host]",
	Short: "Scan all opened TCP or UDP ports.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		module.Log(level) // set logger level

	},
}

func init() {
	rootCmd.AddCommand(portCmd)

	// portCmd
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
