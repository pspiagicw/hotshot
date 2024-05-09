package object

import "os"

func exitFunc(args []Object) Object {
	os.Exit(0)
	return &Null{}
}
