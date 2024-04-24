package main

import (
	"github.com/pspiagicw/hotshot/argparse"
	"github.com/pspiagicw/hotshot/handle"
)

func main() {
	opts := argparse.ParseArguments()
	handle.Handle(opts)
}
