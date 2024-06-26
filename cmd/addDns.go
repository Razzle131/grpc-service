/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	"github.com/Razzle131/grpc-service/internal/consts"
	desc "github.com/Razzle131/grpc-service/internal/generated"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// addDnsCmd represents the addDns command
var addDnsCmd = &cobra.Command{
	Use:   "addDns",
	Short: "adds given ip to dns list",
	Long:  "adds given ip address to /etc/resolv.conf file (resets after reboot)",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(consts.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to server: %v", err)
		}
		defer conn.Close()

		c := desc.NewCustomizerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), waitTime)
		defer cancel()

		_, err = c.AddDNS(ctx, &desc.SetDnsRequest{DnsIp: args[0], SudoPassword: args[1]})
		if err != nil {
			log.Fatalf("failed to add dns: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addDnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addDnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addDnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
