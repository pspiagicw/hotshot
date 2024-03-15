package object

func pushFunc(args []Object) Object {
	err := assertArity("PUSH", args, 2)
	if err != nil {
		return err
	}

	table := args[0]
	if table.Type() != TABLE_OBJ {
		return createError("Object not a table.")
	}

	t, ok := table.(*Table)
	if !ok {
		return createError("Couldn't cast object to table.")
	}

	value := args[1]

	t.Elements = append(t.Elements, value)

	return Null{}
}
func popFunc(args []Object) Object {
	err := assertArity("POP", args, 1)
	if err != nil {
		return err
	}

	table := args[0]
	if table.Type() != TABLE_OBJ {
		return createError("Object not a table.")
	}

	t, ok := table.(*Table)
	if !ok {
		return createError("Couldn't cast object to table.")
	}

	count := len(t.Elements)

	if count == 0 {
		return createError("Attempt to pop from empty table!")
	}

	value := t.Elements[count-1]

	t.Elements = t.Elements[:count-1]
	return value
}
func carFunc(args []Object) Object {
	err := assertArity("CAR", args, 1)
	if err != nil {
		return err
	}

	table := args[0]
	if table.Type() != TABLE_OBJ {
		return createError("Object not a table.")
	}

	t, ok := table.(*Table)
	if !ok {
		return createError("Couldn't cast object to table.")
	}

	count := len(t.Elements)

	if count == 0 {
		return createError("Attempt to car from empty table!")
	}

	return t.Elements[0]
}
func cdrFunc(args []Object) Object {
	err := assertArity("CDR", args, 1)
	if err != nil {
		return err
	}

	table := args[0]
	if table.Type() != TABLE_OBJ {
		return createError("Object not a table.")
	}

	t, ok := table.(*Table)
	if !ok {
		return createError("Couldn't cast object to table.")
	}

	count := len(t.Elements)

	if count == 0 {
		return createError("Attempt to cdr from empty table!")
	}

	return Table{
		Elements: t.Elements[1:],
	}
}
func listFunc(args []Object) Object {
	err := assertArgs("LIST", args)
	if err != nil {
		return err
	}

	return &Table{
		Elements: args,
	}
}
func reverseFunc(args []Object) Object {
	err := assertArgs("REVERSE", args)
	if err != nil {
		return err
	}

	value := args[0].(*Table)

	reverseElements := []Object{}

	length := len(value.Elements)

	for i := length - 1; i >= 0; i-- {
		reverseElements = append(reverseElements, value.Elements[i])
	}
	return &Table{
		Elements: reverseElements,
	}
}
func lastFunc(args []Object) Object {
	err := assertArity("LAST", args, 1)
	if err != nil {
		return err
	}

	table := args[0].(*Table)

	length := len(table.Elements)

	if length == 0 {
		return createError("Attempt to get last value from empty table")
	}

	return table.Elements[length-1]
}
