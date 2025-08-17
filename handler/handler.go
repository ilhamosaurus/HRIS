package handler

import "github.com/ilhamosaurus/HRIS/pkg/util"

type Handler struct {
	Hasher *util.Hasher
}

func NewHandler(hasher *util.Hasher) *Handler {
	return &Handler{Hasher: hasher}
}
