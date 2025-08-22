package main

import (
	"context"
	"log"
	"os"

	"github.com/jacobbrewer1/trackify/username"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "username",
				Usage: "Track online accounts for a given username",
				Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
					args := cmd.Args()
					switch {
					case args.Len() < 1:
						return ctx, cli.Exit("username is required", 1)
					case args.Len() > 1:
						return ctx, cli.Exit("too many arguments", 1)
					}
					return ctx, nil
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return username.TrackUsername(
						ctx,
						cmd.Args().First(),
					)
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
