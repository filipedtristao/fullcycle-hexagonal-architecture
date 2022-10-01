/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/filipedtristao/hexagonal-architecture/adapters/db"
	"github.com/filipedtristao/hexagonal-architecture/adapters/cli"
	"github.com/filipedtristao/hexagonal-architecture/application"
	"github.com/spf13/cobra"
	"database/sql"
	"fmt"
)

var conn, _ = sql.Open("sqlite3", "db.sqlite3")
var productDb = db.NewProductDb(conn)
var productService = application.ProductService{ProductPersistence: productDb}

var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(&productService, action, productId, productName, productPrice)

		if err != nil {
			panic(err)
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "", "Action to be performed")
	cliCmd.Flags().StringVarP(&productId, "product-id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "product-name", "n", "", "Product name")
	cliCmd.Flags().Float64VarP(&productPrice, "product-price", "p", 0, "Product price")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
