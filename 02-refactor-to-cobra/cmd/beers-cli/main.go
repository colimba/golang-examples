package main

import (
	"fmt"
	"os"

	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	beerCmd := &cobra.Command{Use: "beers-cli"}
	storeCmd := &cobra.Command{Use: "stores-cli"}

	beerCmd.AddCommand(cli.InitBeersCmd())
	storeCmd.AddCommand(cli.InitStoresCmd())

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "beers":
			beerCmd.SetArgs(os.Args[1:])
			beerCmd.Execute()
		case "stores":
			storeCmd.SetArgs(os.Args[1:])
			storeCmd.Execute()
		default:
			fmt.Println("Invalid command")
		}

	}

}
