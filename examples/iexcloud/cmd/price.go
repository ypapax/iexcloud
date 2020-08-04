// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/ypapax/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"context"
	"fmt"
	"log"

	iex "github.com/ypapax/iexcloud/v2"
	"github.com/ypapax/iexcloud/v2/examples/iexcloud/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(priceCmd)
}

var priceCmd = &cobra.Command{
	Use:   "price [stock]",
	Short: "Retrieve the current price for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		price, err := client.Price(context.Background(), stock)
		if err != nil {
			log.Fatalf("Error getting stock price: %s", err)
		}
		fmt.Printf("Current price = $%.2f\n", price)
	},
}
