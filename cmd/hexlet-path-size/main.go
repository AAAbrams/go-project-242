package main

import (
	"code/cmd/internal"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		UsageText: "hexlet-path-size [options] <path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "path",
				UsageText: "path to file or dir",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			path := c.StringArg("path")
			humanView := c.Bool("human")
			scanAll := c.Bool("all")
			scanRecursive := c.Bool("recursive")

			if path != "" {
				size, err := internal.GetPathSize(path, scanAll, scanRecursive)
				if err != nil {
					return err
				}
				output := internal.OutputFmt(size, path, humanView)
				fmt.Print(output)
			} else {
				err := cli.ShowAppHelp(c)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
