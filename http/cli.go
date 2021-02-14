package http

import "github.com/urfave/cli"

// HTTPCommand is used for CLI entry
var HTTPCommand = cli.Command{
	Name:   "http",
	Usage:  "generate an http filesystem",
	Action: httpAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "package",
			Value: packageName,
		},
		cli.StringFlag{
			Name:  "input",
			Value: defaultPattern,
		},
		cli.StringFlag{
			Name:  "output",
			Usage: "defines the output destionation. Must match: {filename}.go",
			// default value
			Value: outputName,
		},
		cli.StringFlag{
			Name:  "trim-prefix",
			Value: "files",
		},
	},
}

func httpAction(c *cli.Context) error {
	return HttpHandle(
		c.String("input"),
		c.String("package"),
		c.String("output"),
		c.String("trim-prefix"),
	)
}
