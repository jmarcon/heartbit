/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hb",
	Short: "Heartbit is focused on create an dotnet environment and projects",
	Long: `Heartbit will be responsible for validate your environment and ensure
that all the requirements are available for dotnet development, environment
variables configured, and tools for development are available

This cli will help you create a dotnet project in a very opnionated
configuration`,
	Version: "0.0.1",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	  },
	  PreRun: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	  },
	  Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Inside rootCmd Run with args: %v\n", args)
	  },
	  PostRun: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	  },
	  PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	  },
	  DisableAutoGenTag: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func GetCommand() **cobra.Command {
	return &rootCmd
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Fatal error")
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hb.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


