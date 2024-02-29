package object

type Environment struct {
	Vars      map[string]Object
	Functions map[string]Object
}

func NewEnvironment() *Environment {
	return &Environment{
		Vars:      map[string]Object{},
		Functions: getBuiltins(),
	}
}
