package models

type FileTypes string

const (
	Main      FileTypes = "main"
	Variables FileTypes = "variables"
	Outputs   FileTypes = "outputs"
)

type File struct {
	Name    string
	Type    FileTypes
	Content []byte
}
