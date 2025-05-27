package main

import (
	"github.com/joel2santos/fabric/cmd"
	"github.com/joel2santos/fabric/cmd/fabric"
)

func main() {
	cmd.RootCmd.AddCommand(
		fabric.FabricEntityCmd,
		fabric.FabricModelCmd,
	)
	cmd.Execute()
}
