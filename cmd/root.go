package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var level string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golire [Command] [flags]",
	Short: "scan tool",
	Long: `
┌─┐┌─┐┬  ┬┬─┐┌─┐
│ ┬│ ││  │├┬┘├┤ 
└─┘└─┘┴─┘┴┴└─└─┘

Author: Secriy
GitHub: https://github.com/secriy/golire
`,
	Version: "0.2",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Long)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&level, "level", "l", "", "logger level, debug, info, warning, error and fatal")
}
