package object

func pushFunc(args []Object) Object {
	if err := assertArity("PUSH", args, 2); err != nil {
		return err
	}

	if err := validateTable(args[0]); err != nil {
		return err
	}

	t := args[0].(*Table)
	value := args[1]

	addToTable(t, value)

	return Null{}
}
func popFunc(args []Object) Object {
	if err := assertArity("POP", args, 1); err != nil {
		return err
	}

	if err := validateTable(args[0]); err != nil {
		return err
	}

	t := args[0].(*Table)

	return popFromTable(t)
}

func carFunc(args []Object) Object {
	if err := assertArity("CAR", args, 1); err != nil {
		return err
	}

	if err := validateTable(args[0]); err != nil {
		return err
	}

	t := args[0].(*Table)

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
	if err := assertArgs("REVERSE", args); err != nil {
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
func addToTable(t *Table, value Object) {
	t.Elements = append(t.Elements, value)
}
func validateTable(object Object) Object {
	if object.Type() != TABLE_OBJ {
		return createError("Object not a table.")
	}
	return nil
}
func popFromTable(t *Table) Object {
	count := len(t.Elements)

	if count == 0 {
		return createError("Attempt to pop from empty table!")
	}

	value := t.Elements[count-1]

	t.Elements = t.Elements[:count-1]
	return value
}
