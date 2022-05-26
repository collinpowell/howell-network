package app

import (
	"os"

	stdlog "log"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

//const appName = "HowellApp"

var (
	// DefaultNodeHome defines default home directories for howelld
	DefaultNodeHome string

	// ModuleBasics = The ModuleBasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
	)

)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		stdlog.Println("Failed to get home dir %2", err)
	}
	
	DefaultNodeHome = filepath.Join(userHomeDir, ".howell")
}
