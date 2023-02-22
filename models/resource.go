package models

/*
resources:

expects 2 labels (type, name): `resources "foo" "bar" {}`


*/

type Resource struct {
	Type string
	Name string

	Arguments []Argument
}
