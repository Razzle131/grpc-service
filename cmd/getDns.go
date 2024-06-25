/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"time"

	"github.com/Razzle131/grpc-service/internal/consts"
	desc "github.com/Razzle131/grpc-service/internal/generated"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// getDnsCmd represents the getDns command
var getDnsCmd = &cobra.Command{
	Use:   "getDns",
	Short: "gets list of current dns ip addresses",
	Long:  "gets list of current dns ip addresses from /etc/resolv.conf file",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(consts.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to server: %v", err)
		}
		defer conn.Close()

		c := desc.NewCustomizerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := c.GetDNS(ctx, &desc.GetDnsRequest{})
		if err != nil {
			log.Fatalf("failed to get note by id: %v", err)
		}

		for _, val := range res.DnsIps {
			log.Println(val)
		}
	},
}

func init() {
	rootCmd.AddCommand(getDnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
