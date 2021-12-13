package cli

import (
	"flag"

	"dhemery.com/panelgen/internal/builder"
)

func Execute() error {
	allFlag := flag.Bool("all", false, "build all modules")
	flag.Parse()
	if *allFlag {
		return builder.BuildAll()
	}
	if flag.NArg() > 0 {
		return builder.Build(flag.Args())
	}
	builder.List()
	return nil
}
