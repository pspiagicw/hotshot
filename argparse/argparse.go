package argparse

import "flag"
import "github.com/pspiagicw/hotshot/interpreter"

type Options struct {
	Debug bool
	Args  []string
}

func ParseArguments() *Options {
	o := new(Options)

	flag.BoolVar(&o.Debug, "debug", false, "Enable debug mode.")

	flag.Parse()

	o.Args = flag.Args()

	return o
}

func HandleOpts(opts *Options) {
	if len(opts.Args) == 0 {
		interpreter.StartREPL()
	} else {
		for _, arg := range opts.Args {
			interpreter.ExecuteFile(arg, opts.Debug)
		}
	}
}
