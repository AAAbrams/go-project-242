package main

import (
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
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "path",
				UsageText: "path to file or dir",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			path := c.StringArg("path")
			if path != "" {
				size, err := GetPathSize(path)
				if err != nil {
					return err
				}
				fmt.Printf("%dB\t%s", size, path)
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
