package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/black-gato/er-cli/cmd"
)

type Globals struct {
	Format string `help:"Output format" default:"console" enum:"console,json"`
}

const appName = "err-cli"

var version string

type VersionFlag string

func (v VersionFlag) Decode(_ *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                       { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type CLI struct {
	Version VersionFlag `       help:"Print version information and quit" short:"v" name:"version"`
	Init    cmd.Init    `cmd:"" help:"Setup configuration for errors content"`
	Add     cmd.Add     `cmd:"" help:"Used to add some type of new entry"`
}

func run() error {
	if version == "" {
		version = "development"
	}
	cli := CLI{
		Version: VersionFlag(version),
	}
	// Display help if no args are provided instead of an error message
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}

	ctx := kong.Parse(&cli,
		kong.Name(appName),
		kong.Description("Search and store Error messages to build a knowledge hub for homegrown tools"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Tree: true,
		}),
		kong.DefaultEnvars(appName),
		kong.Vars{
			"version": string(cli.Version),
		})
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
	return nil
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
