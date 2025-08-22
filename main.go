package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/jacobbrewer1/trackify/username"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "username",
				Usage: "Track online accounts for a given username",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "target",
						Aliases: []string{"t"},
						Usage:   "Specify a target platform to search (can be specified multiple times). If not specified, all platforms will be searched.",
					},
				},
				Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
					args := cmd.Args()
					if args.Len() < 1 {
						return nil, cli.Exit("at least one username is required", 1)
					}
					return ctx, nil
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return username.TrackUsernames(
						ctx,
						cmd.Args().Slice(),
						cmd.StringSlice("target"),
					)
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
