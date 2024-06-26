/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/Razzle131/grpc-service/internal/consts"
	desc "github.com/Razzle131/grpc-service/internal/generated"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// setHostnameCmd represents the setHostname command
var setHostnameCmd = &cobra.Command{
	Use:   "setHostname",
	Short: "sets servers hostname with given arg",
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

		res, err := c.SetHostName(ctx, &desc.SetHostRequest{NewHostname: args[0], SudoPassword: args[1]})
		if err != nil {
			log.Fatalf("failed to set hostname: %v", err)
		}

		fmt.Println(res.CurHostname)
	},
}

func init() {
	rootCmd.AddCommand(setHostnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setHostnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setHostnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
