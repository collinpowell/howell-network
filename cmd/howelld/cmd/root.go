package cmd

import (
	howellapp "github.com/collinpowell/howell-network/app"
	"github.com/collinpowell/howell-network/app/params"
	"github.com/spf13/cobra"

	"fmt"
)

func NewRootCmd() (*cobra.Command, params.EncodingConfig) {

	// Encoding Configurations for Codec Information Transfer
	encodingConfig := howellapp.MakeEncodingConfig()


	fmt.Print(encodingConfig)

	rootCmd := &cobra.Command{
		Use:   "howelld",
		Short: "Eminent Howell App",
	}

	return rootCmd,encodingConfig

}