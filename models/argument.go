package models

const (
	VARIABLE = iota
	DATA
	HARDCODED
)

type Argument struct {
	Type int

	Key            string
	ReferenceValue string
	Value          string

	CustomType string
	Hardcoded  string

	Edit bool

	Object []Argument
}
