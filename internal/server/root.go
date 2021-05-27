package server

import "github.com/spf13/cobra"

var rootCommand = cobra.Command{
	Use:  "pennyPincher",
	Long: "Finds best buy price for BTC across multiple exchanges",
}

func Execute() {
	rootCommand.Execute()
}
