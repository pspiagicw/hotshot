package object

func registerBuiltin(builtin map[string]*Builtin, name string, builtinFunc BuiltinFunc) {
	builtin[name] = &Builtin{
		Fn: builtinFunc,
	}

}

type BuiltinIndex struct {
	Name string
	Func *Builtin
}

func BuiltinList() []BuiltinIndex {
	list := []BuiltinIndex{}

	builtins := getBuiltins()

	list = appendBuiltin(list, "echo", builtins) // 0
	list = appendBuiltin(list, "len", builtins)
	list = appendBuiltin(list, "do", builtins)

	list = appendBuiltin(list, "not", builtins)
	list = appendBuiltin(list, "and", builtins)
	list = appendBuiltin(list, "or", builtins)

	list = appendBuiltin(list, "numberp", builtins)
	list = appendBuiltin(list, "stringp", builtins)
	list = appendBuiltin(list, "tablep", builtins)
	list = appendBuiltin(list, "functionp", builtins)

	list = appendBuiltin(list, "mod", builtins)
	list = appendBuiltin(list, "push", builtins)
	list = appendBuiltin(list, "cons", builtins)
	list = appendBuiltin(list, "pop", builtins)
	list = appendBuiltin(list, "car", builtins)
	list = appendBuiltin(list, "cdr", builtins)
	list = appendBuiltin(list, "list", builtins)
	list = appendBuiltin(list, "extend", builtins)
	list = appendBuiltin(list, "reverse", builtins)

	list = appendBuiltin(list, "sqrt", builtins)
	list = appendBuiltin(list, "exp", builtins)
	list = appendBuiltin(list, "min", builtins)
	list = appendBuiltin(list, "max", builtins)

	list = appendBuiltin(list, "concat", builtins)
	list = appendBuiltin(list, "string", builtins)
	list = appendBuiltin(list, "last", builtins)
	list = appendBuiltin(list, "getchar", builtins)
	list = appendBuiltin(list, "substring", builtins)
	list = appendBuiltin(list, "lowercase", builtins)
	list = appendBuiltin(list, "uppercase", builtins)

	list = appendBuiltin(list, "count", builtins)
	list = appendBuiltin(list, "type", builtins)
	list = appendBuiltin(list, "exit", builtins)
	list = appendBuiltin(list, "inc", builtins)
	list = appendBuiltin(list, "dec", builtins)
	list = appendBuiltin(list, "input", builtins)
	list = appendBuiltin(list, "return", builtins)

	list = appendBuiltin(list, "+", builtins)
	list = appendBuiltin(list, "-", builtins)
	list = appendBuiltin(list, "*", builtins)
	list = appendBuiltin(list, "/", builtins)
	list = appendBuiltin(list, "=", builtins)
	list = appendBuiltin(list, "<", builtins)
	list = appendBuiltin(list, ">", builtins)

	return list
}
func appendBuiltin(list []BuiltinIndex, name string, table map[string]*Builtin) []BuiltinIndex {
	return append(list, BuiltinIndex{
		Name: name,
		Func: table[name],
	})
}
func Essentials() map[string]*Builtin {
	builtins := map[string]*Builtin{}

	registerBuiltin(builtins, "+", addFunc)
	registerBuiltin(builtins, "-", minusFunc)
	registerBuiltin(builtins, "*", multiplyFunc)
	registerBuiltin(builtins, "/", divideFunc)
	registerBuiltin(builtins, "=", equalFunc)
	registerBuiltin(builtins, "<", ltFunc)
	registerBuiltin(builtins, ">", gtFunc)

	return builtins
}

func getBuiltins() map[string]*Builtin {
	builtins := map[string]*Builtin{}

	registerBuiltin(builtins, "+", addFunc)
	registerBuiltin(builtins, "-", minusFunc)
	registerBuiltin(builtins, "*", multiplyFunc)
	registerBuiltin(builtins, "/", divideFunc)
	registerBuiltin(builtins, "=", equalFunc)
	registerBuiltin(builtins, "<", ltFunc)
	registerBuiltin(builtins, ">", gtFunc)

	registerBuiltin(builtins, "echo", printFunc)
	registerBuiltin(builtins, "len", lenFunc)
	registerBuiltin(builtins, "do", doFunc)

	registerBuiltin(builtins, "not", notFunc)
	registerBuiltin(builtins, "and", andFunc)
	registerBuiltin(builtins, "or", orFunc)

	registerBuiltin(builtins, "numberp", numberpFunc)
	registerBuiltin(builtins, "stringp", stringpFunc)
	registerBuiltin(builtins, "tablep", tablepFunc)
	registerBuiltin(builtins, "functionp", functionpFunc)

	registerBuiltin(builtins, "mod", modFunc)
	registerBuiltin(builtins, "push", pushFunc)
	registerBuiltin(builtins, "cons", pushFunc)
	registerBuiltin(builtins, "pop", popFunc)
	registerBuiltin(builtins, "car", carFunc)
	registerBuiltin(builtins, "cdr", cdrFunc)
	registerBuiltin(builtins, "list", listFunc)
	registerBuiltin(builtins, "extend", extendFunc)
	registerBuiltin(builtins, "reverse", reverseFunc)

	registerBuiltin(builtins, "sqrt", sqrtFunc)
	registerBuiltin(builtins, "exp", expFunc)
	registerBuiltin(builtins, "min", minFunc)
	registerBuiltin(builtins, "max", maxFunc)

	registerBuiltin(builtins, "concat", concatFunc)
	registerBuiltin(builtins, "string", stringFunc)
	registerBuiltin(builtins, "last", lastFunc)
	registerBuiltin(builtins, "getchar", getCharFunc)
	registerBuiltin(builtins, "substring", subStringFunc)
	registerBuiltin(builtins, "lowercase", lowerFunc)
	registerBuiltin(builtins, "uppercase", upperFunc)

	registerBuiltin(builtins, "count", countFunc)
	registerBuiltin(builtins, "type", typeFunc)
	registerBuiltin(builtins, "exit", exitFunc)
	registerBuiltin(builtins, "inc", incFunc)
	registerBuiltin(builtins, "dec", decFunc)
	registerBuiltin(builtins, "input", inputFunc)
	registerBuiltin(builtins, "return", returnFunc)

	return builtins
}
