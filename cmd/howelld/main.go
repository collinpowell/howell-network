package main

import (
	"os"

	howellapp "github.com/collinpowell/howell-network/app"
	"github.com/collinpowell/howell-network/cmd/howelld/cmd"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main(){
	rootCmd,_ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, howellapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			//os.Exit(1)
		}
	}
}