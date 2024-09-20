package main

import (
	"fmt"
	"os"


	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/kiarash-naderi/myapp/app"
	"github.com/kiarash-naderi/myapp/cmd/myappd/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "MYAPP", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
