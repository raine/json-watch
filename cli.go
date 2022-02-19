package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

type Options struct {
	help bool
	name string
	key  string
}

var f *flag.FlagSet

func parseArgs(args []string) (Options, error) {
	opts := Options{}
	f = flag.NewFlagSet("", flag.ContinueOnError)
	f.BoolVarP(&opts.help, "help", "h", false, "show this help")
	f.StringVarP(&opts.key, "key", "k", "", "prop in json objects that identifies them (basically the id)")
	err := f.Parse(args[1:])
	opts.name = f.Arg(0)

	if opts.help {
		printUsage()
		os.Exit(0)
	}

	if opts.name == "" {
		return opts, fmt.Errorf("name is required")
	}

	if opts.key == "" {
		return opts, fmt.Errorf("key is required")
	}

	return opts, err
}

func printUsage() {
	usage := `Usage: cat data.json | json-watch <name>

Takes a list of objects as JSON through stdin.

The first execution will "prime the watch file" and won't print output.

On further executions, unseen JSON objects in the array will be printed to
stdout as newline delimited JSON.

The name parameter uniquely identifies an instance of json-watch usage, so if
you are watching multiple JSONs for new objects, each of the json-watch calls
should have a distinct name.

Options:`
	fmt.Println(usage)
	f.PrintDefaults()
}
