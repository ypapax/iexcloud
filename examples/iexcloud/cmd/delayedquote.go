// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/ypapax/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	iex "github.com/ypapax/iexcloud/v2"
	"github.com/ypapax/iexcloud/v2/examples/iexcloud/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(delayedQuoteCmd)
}

var delayedQuoteCmd = &cobra.Command{
	Use:   "dq [stock]",
	Short: "Retrieve the 15 minute delayed quote for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		dq, err := client.DelayedQuote(context.Background(), stock)
		b, err := json.MarshalIndent(dq, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Delayed Quote ##")
		fmt.Println(string(b))
	},
}
