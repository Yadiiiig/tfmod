package utils

import (
	"os"
	"strings"

	"github.com/yadiiiig/tfmod/models"
)

func OpenFile(name string) (*models.File, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return &models.File{
		Name:    name,
		Type:    models.FileTypes(strings.Split(name, ".")[0]),
		Content: data,
	}, nil
}
