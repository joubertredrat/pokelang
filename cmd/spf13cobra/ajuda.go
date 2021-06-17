package spf13cobra

import (
	"github.com/joubertredrat/pokelang/cmd"
	"github.com/spf13/cobra"
)

var ajuda = &cobra.Command{
	Use:   "ajuda",
	Short: "Exibe a ajuda de como jogar",
	Run: func(c *cobra.Command, args []string) {
		cmd.ShowHelp()
	},
}
