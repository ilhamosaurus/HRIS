package handler

import (
	"github.com/ilhamosaurus/HRIS/model"
	"github.com/ilhamosaurus/HRIS/pkg/util"
)

type Handler struct {
	hasher *util.Hasher
	model  *model.Model
}

func NewHandler(hasher *util.Hasher, model *model.Model) *Handler {
	return &Handler{
		hasher: hasher,
		model:  model,
	}
}
