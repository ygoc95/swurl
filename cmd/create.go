package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ygoc95/swurl/services"
)

type Endpoint struct {
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates .hurl files by taking swagger.json url or directly json file itself.",
	Long:  `Creates .hurl files by taking swagger.json url or directly json file itself.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating test cases")
		url, _ := cmd.Flags().GetString("url")
		services.CreateHurlFile(url)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("url", "u", "", "URL for swagger.json")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}