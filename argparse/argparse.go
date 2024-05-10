package argparse

import (
	"flag"

	"github.com/pspiagicw/hotshot/help"
)

type Opts struct {
	AST  bool
	Args []string

	Token bool
	Null  bool
}

func ParseArguments() *Opts {
	o := new(Opts)
	flag.Usage = help.Help

	flag.BoolVar(&o.AST, "print-ast", false, "Print AST")
	flag.BoolVar(&o.Token, "print-tokens", false, "Print Tokens")
	flag.BoolVar(&o.Null, "null", false, "Print Null in REPL")

	flag.Parse()

	o.Args = flag.Args()

	return o
}
