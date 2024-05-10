package object

func registerBuiltin(builtin map[string]*Builtin, name string, builtinFunc BuiltinFunc) {
	builtin[name] = &Builtin{
		Fn: builtinFunc,
	}

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

	registerBuiltin(builtins, "not", notFunc)
	registerBuiltin(builtins, "and", andFunc)
	registerBuiltin(builtins, "or", orFunc)

	registerBuiltin(builtins, "do", doFunc)
	registerBuiltin(builtins, "echo", printFunc)
	registerBuiltin(builtins, "len", lenFunc)

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
