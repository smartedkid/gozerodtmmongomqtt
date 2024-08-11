/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"2112a-6/month/wsyx/cobra/consumer"
	"github.com/spf13/cobra"
)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Order_Stock,
}

func Order_Stock(cmd *cobra.Command, args []string) {
	consumer.Consumer()
}
func init() {
	rootCmd.AddCommand(stockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
