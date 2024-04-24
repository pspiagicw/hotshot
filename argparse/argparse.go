package argparse

import "flag"

type Opts struct {
	AST  bool
	Args []string

	Token bool
}

func ParseArguments() *Opts {
	o := new(Opts)

	flag.BoolVar(&o.AST, "print-ast", false, "Print AST")
	flag.BoolVar(&o.Token, "print-tokens", false, "Print Tokens")

	flag.Parse()

	o.Args = flag.Args()

	return o
}
