package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "fabric",
	Short: "Fabric is a tool for generating code based on entity and model definitions.",
	Long: "Fabric is a command-line tool that helps developers generate code based on entity and model definitions",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
