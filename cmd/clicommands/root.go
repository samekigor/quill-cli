package clicommands

import (
	"log"
	"os"

	"github.com/samekigor/quill-cli/cmd/clicommands/auths"
	"github.com/samekigor/quill-cli/cmd/client"
	"github.com/samekigor/quill-cli/cmd/internal/utils"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "quill",
		Short: "Lightweight container system with flexible configuration management",
		Long:  `Quill is a CLI tool that allows users to manage and deploy containerized applications with ease.`,
	}
)

func Execute() {
	utils.InitEnviromentVariables()
	defer func() {
		if client.GrpcClient != nil {
			client.GrpcClient.Close()
		}
	}()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initClient)
	rootCmd.AddCommand(auths.LoginCmd)
	rootCmd.AddCommand(auths.LogoutCmd)
}

func initClient() {
	var err error
	client.SocketPath, err = utils.GetEnviromentVariable("SOCKET_PATH")
	if err != nil {
		log.Fatalf("QUILL_SOCKET_PATH enviroment variable not found: %v", err)
	}
	client.GrpcClient, err = client.NewGRPCClient(client.SocketPath)
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
}
