package code

//go:generate stringer -type=Op

type Op int

const (
	RETURN Op = iota
	PUSH
	DIV
	ADD
	MUL
	SUB

	TRUE
	FALSE

	GT
	LT
	EQ

	JCMP
	JMP
	SET
	GET
	LSET
	LGET

	JT
	CALL
	BUILTIN

	TABLE
	ASSERT
	INDEX
	DICT
)

type Instruction struct {
	OpCode  Op
	Operand int16
	Comment string
}
