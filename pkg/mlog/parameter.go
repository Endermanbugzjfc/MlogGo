package mlog

type CodeBlockParameter struct {
	expectedTypes []DataType
	value         Variable
}

type Variable struct {
	// TODO: Variable.
}

type DataType interface {
	Validate(Variable) bool
}
