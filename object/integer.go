package object

import (
	"math"
)

func incFunc(args []Object) Object {
	err, values := assertArgsInteger("INC", args)
	if err != nil {
		return err
	}

	for _, val := range values {
		val.Value++
	}

	return &Null{}
}
func decFunc(args []Object) Object {
	err, values := assertArgsInteger("DEC", args)
	if err != nil {
		return err
	}

	for _, val := range values {
		val.Value--
	}

	return &Null{}
}

func multiplyFunc(args []Object) Object {
	err, values := assertArgsInteger("MULTIPLY", args)
	if err != nil {
		return err
	}

	result := 1
	for _, val := range values {
		result *= val.Value
	}

	return &Integer{
		Value: result,
	}

}
func divideFunc(args []Object) Object {

	err, values := assertArgsInteger("DIVIDE", args)
	if err != nil {
		return err
	}

	first := values[0]
	result := first.Value
	for _, val := range values[1:] {
		result /= val.Value
	}

	return &Integer{
		Value: result,
	}
}
func minusFunc(args []Object) Object {

	err, values := assertArgsInteger("MINUS", args)
	if err != nil {
		return err
	}

	result := values[0].Value
	for _, val := range values[1:] {
		result -= val.Value
	}

	return &Integer{
		Value: result,
	}
}

func addFunc(args []Object) Object {

	err, values := assertArgsInteger("ADD", args)
	if err != nil {
		return err
	}

	result := 0
	for _, val := range values {
		result += val.Value
	}

	return &Integer{
		Value: result,
	}

}
func gtFunc(args []Object) Object {
	err, values := assertArgsInteger("GT", args)

	if err != nil {
		return err
	}

	cmp := math.MaxInt
	result := true
	for _, val := range values {
		result = result && cmp > val.Value
		cmp = val.Value
	}

	return &Boolean{
		Value: result,
	}

}
func ltFunc(args []Object) Object {
	err, values := assertArgsInteger("LT", args)

	if err != nil {
		return err
	}

	cmp := math.MinInt
	result := true
	for _, val := range values {
		result = result && cmp < val.Value
		cmp = val.Value
	}

	return &Boolean{
		Value: result,
	}
}
func equalFunc(args []Object) Object {
	err := assertArgs("EQUAL", args)
	if err != nil {
		return err
	}

	result := true
	var cmp Object = args[0]
	for _, val := range args {
		result = result && isObjectEqual(val, cmp)
		cmp = val
	}

	return &Boolean{
		Value: result,
	}

}
func modFunc(args []Object) Object {
	err, values := assertArgsInteger("MOD", args)
	if err != nil {
		return err
	}

	result := values[0].Value
	for _, val := range values[1:] {
		result = result % val.Value
	}

	return &Integer{
		Value: result,
	}
}
func sqrtFunc(args []Object) Object {
	err := assertArity("SQRT", args, 1)
	if err != nil {
		return err
	}

	value, ok := args[0].(*Integer)
	if !ok {
		return createError("SQRT expects Integer, found %v", args[0].Type())
	}

	return &Integer{
		Value: int(math.Sqrt(float64(value.Value))),
	}
}
func expFunc(args []Object) Object {
	err := assertArity("EXP", args, 2)
	if err != nil {
		return err
	}

	base, ok := args[0].(*Integer)
	if !ok {
		return createError("EXP expects Integer, found %v", args[0].Type())
	}

	exp, ok := args[1].(*Integer)
	if !ok {
		return createError("EXP expects Integer, found %v", args[1].Type())
	}

	return &Integer{
		Value: int(math.Pow(float64(base.Value), float64(exp.Value))),
	}
}
func minFunc(args []Object) Object {
	err, values := assertArgsInteger("MIN", args)
	if err != nil {
		return err
	}

	min := &Integer{Value: math.MaxInt}
	for _, val := range values {

		if val.Value < min.Value {
			min = val
		}
	}

	return min
}
func maxFunc(args []Object) Object {
	err, values := assertArgsInteger("MAX", args)
	if err != nil {
		return err
	}

	max := &Integer{Value: math.MinInt}
	for _, val := range values {
		if val.Value > max.Value {
			max = val
		}
	}

	return max
}
