package main

import (
	"fmt"
	"os"

	"github.com/joubertredrat/pokelang/cmd/purist"
	"github.com/joubertredrat/pokelang/cmd/spf13cobra"
	"github.com/joubertredrat/pokelang/cmd/urfavecli"
)

func main() {

	switch os.Getenv("CMD_TYPE") {
	case "purist":
		purist.Execute()
	case "spf13cobra":
		spf13cobra.Execute()
	case "urfavecli":
		urfavecli.Execute()
	default:
		fmt.Println("CMD_TYPE env not defined, available options: purist, spf13cobra, urfavecli")
	}

}
