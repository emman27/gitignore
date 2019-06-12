package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emman27/gitignore/pkg/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	rootCmd := &cobra.Command{
		Use:   "gitignore",
		Short: "Gitignore generator",
	}
	rootCmd.AddCommand(cli.GenerateCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
