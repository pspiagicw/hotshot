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
)

type Instruction struct {
	OpCode  Op
	Args    int16
	Comment string
}
