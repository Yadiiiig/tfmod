package models

type Argument struct {
	Key   string
	Value string

	Type       int
	CustomType string
	Hardcoded  string

	Edit bool

	Childs []Argument
}
