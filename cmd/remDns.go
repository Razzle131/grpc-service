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

// remDnsCmd represents the remDns command
var remDnsCmd = &cobra.Command{
	Use:   "remDns",
	Short: "removes given dns ip address",
	Long:  "removes given dns ip address from /etc/resolv.conf file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(consts.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to server: %v", err)
		}
		defer conn.Close()

		c := desc.NewCustomizerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err = c.RemoveDNS(ctx, &desc.SetDnsRequest{DnsIp: args[0]})
		if err != nil {
			log.Fatalf("failed to get note by id: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(remDnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remDnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remDnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}