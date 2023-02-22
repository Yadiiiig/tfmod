package models

type Field struct {
	Key string

	Type       int
	CustomType string
	Hardcoded  string

	Data string

	Edit bool

	Childs []Field
}
