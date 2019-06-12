// Package cli contains a CLI controller to the gitignore workflow
package cli

import (
	"context"
	"fmt"

	"github.com/emman27/gitignore/pkg/gitignore"
	"github.com/emman27/gitignore/pkg/output"
	"github.com/spf13/cobra"
)

var cliInstance *cli

func getCLIInstance() *cli {
	if cliInstance == nil {
		svc, err := gitignore.NewService()
		if err != nil {
			panic(err)
		}
		output, err := output.New()
		if err != nil {
			panic(err)
		}
		cliInstance = &cli{
			gitignore: svc,
			output:    output,
		}
	}
	return cliInstance
}

type cli struct {
	gitignore gitignore.Service
	output    output.Output
}

func (c *cli) Generate(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("argument length incorrect, expected 1 but got %d", len(args))
	}
	gitignoreType := args[0]
	ctx := context.TODO()
	res, err := c.gitignore.Generate(ctx, gitignoreType)
	if err != nil {
		return err
	}
	err = c.output.Write(ctx, res)
	return err
}

var GenerateCommand = &cobra.Command{
	Use:     "generate <language>",
	Short:   "Generate a gitignore.",
	Example: "gitignore generate go",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := getCLIInstance()
		err := c.Generate(args)
		if err != nil {
			return err
		}
		return nil
	},
}
