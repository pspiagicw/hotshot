package object

type Environment struct {
	Bindings map[string]Object
	Builtins map[string]*Builtin
}

func NewEnvironment() *Environment {
	return &Environment{
		Bindings: map[string]Object{},
		Builtins: getBuiltins(),
	}
}

func (e *Environment) Get(name string) Object {
	value, ok := e.Bindings[name]
	if !ok {
		return createError("No such variable or function!")
	}

	return value
}
func (e *Environment) Set(name string, value Object) {
	e.Bindings[name] = value
}
func (e *Environment) GetBuiltin(name string) Object {
	fn, ok := e.Builtins[name]

	if !ok {
		return nil
	}

	return fn
}
