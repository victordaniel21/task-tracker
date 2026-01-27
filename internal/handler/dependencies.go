package handler

import (
	"github.com/victordaniel21/task-tracker/internal/data"
)

type Dependencies struct {
	Models data.Models
}

func NewDependencies(models data.Models) *Dependencies {
	return &Dependencies{
		Models: models,
	}
}
