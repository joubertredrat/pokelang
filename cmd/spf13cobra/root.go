package spf13cobra

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func Execute() {
	filename := filepath.Base(os.Args[0])

	rootCmd := &cobra.Command{
		Use:   filename,
		Short: "Pokelang: Bora caçar pokemons com golang?",
	}

	rootCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
	rootCmd.AddCommand(NewCapturarCommand())
	rootCmd.AddCommand(ajuda)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(nil, "Error to execute cobra command", err)
	}
}
