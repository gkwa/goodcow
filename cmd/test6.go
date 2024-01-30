/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/goodcow/embedded6"
)

// test6Cmd represents the test6 command
var test6Cmd = &cobra.Command{
	Use:   "test6",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test6 called")
		err := embedded6.RunEmbedded()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(test6Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test6Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test6Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
