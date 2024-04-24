package handle

import (
	"github.com/pspiagicw/hotshot/argparse"
	"github.com/pspiagicw/hotshot/interpreter"
)

func Handle(opts *argparse.Opts) {
	if len(opts.Args) == 0 {
		interpreter.StartREPL(opts)
	} else {
		for _, arg := range opts.Args {
			interpreter.ExecuteFile(arg, opts)
		}
	}

}
