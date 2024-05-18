package main

import (
	"github.com/alecthomas/kong"
)

// will be overwritten in release pipeline
var version = "dev"

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("gstring"),
		kong.Description("A tool to work with strings"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}))
	err := ctx.Run(&cli.Globals)
	if err != nil {
		panic(err)
	}
}
