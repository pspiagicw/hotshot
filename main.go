package main

import (
	"github.com/pspiagicw/hotshot/argparse"
)

func main() {
	opts := argparse.ParseArguments()
	argparse.HandleOpts(opts)
}
