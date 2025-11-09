package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/git"
)

const (
	flagRepository    = "repository"
	flagBranch        = "branch"
	flagRemote        = "remote"
	flagCommitMessage = "message"
	flagAuthor        = "author"
	flagEmail         = "email"
	flagForce         = "force"
	flagVerbose       = "verbose"
	flagQuiet         = "quiet"
	flagDryRun        = "dry-run"
	flagAll           = "all"
	flagInteractive   = "interactive"
)

func main() {
	app := &cli.App{
		Name:  "git",
		Usage: "git command wrapper for yupsh",
		UsageText: `git [OPTIONS] COMMAND [ARG...]

   Execute git commands.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  flagRepository,
				Usage: "repository path or URL",
			},
			&cli.StringFlag{
				Name:    flagBranch,
				Aliases: []string{"b"},
				Usage:   "branch name",
			},
			&cli.StringFlag{
				Name:  flagRemote,
				Usage: "remote name",
			},
			&cli.StringFlag{
				Name:    flagCommitMessage,
				Aliases: []string{"m"},
				Usage:   "commit message",
			},
			&cli.StringFlag{
				Name:  flagAuthor,
				Usage: "author name",
			},
			&cli.StringFlag{
				Name:  flagEmail,
				Usage: "author email",
			},
			&cli.BoolFlag{
				Name:    flagForce,
				Aliases: []string{"f"},
				Usage:   "force operation",
			},
			&cli.BoolFlag{
				Name:    flagVerbose,
				Aliases: []string{"v"},
				Usage:   "be verbose",
			},
			&cli.BoolFlag{
				Name:    flagQuiet,
				Aliases: []string{"q"},
				Usage:   "be quiet",
			},
			&cli.BoolFlag{
				Name:    flagDryRun,
				Aliases: []string{"n"},
				Usage:   "dry run",
			},
			&cli.BoolFlag{
				Name:    flagAll,
				Aliases: []string{"a"},
				Usage:   "all files",
			},
			&cli.BoolFlag{
				Name:    flagInteractive,
				Aliases: []string{"i"},
				Usage:   "interactive mode",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "git: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.IsSet(flagRepository) {
		params = append(params, Repository(c.String(flagRepository)))
	}
	if c.IsSet(flagBranch) {
		params = append(params, Branch(c.String(flagBranch)))
	}
	if c.IsSet(flagRemote) {
		params = append(params, Remote(c.String(flagRemote)))
	}
	if c.IsSet(flagCommitMessage) {
		params = append(params, CommitMessage(c.String(flagCommitMessage)))
	}
	if c.IsSet(flagAuthor) {
		params = append(params, Author(c.String(flagAuthor)))
	}
	if c.IsSet(flagEmail) {
		params = append(params, Email(c.String(flagEmail)))
	}
	if c.Bool(flagForce) {
		params = append(params, Force)
	}
	if c.Bool(flagVerbose) {
		params = append(params, Verbose)
	}
	if c.Bool(flagQuiet) {
		params = append(params, Quiet)
	}
	if c.Bool(flagDryRun) {
		params = append(params, DryRun)
	}
	if c.Bool(flagAll) {
		params = append(params, All)
	}
	if c.Bool(flagInteractive) {
		params = append(params, Interactive)
	}

	// Create and execute the git command
	cmd := Git(params...)
	return gloo.Run(cmd)
}
