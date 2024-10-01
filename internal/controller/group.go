package controller

import (
	"log/slog"

	transfer "github.com/utf2/utf-account-service/internal/controller/model"
)

type GroupControllerAPI interface {
	Search(transfer.GroupSearchRequest) transfer.GroupSearchResponse
}

type GroupController struct {
	log *slog.Logger
}

func NewGroupController(log *slog.Logger) *GroupController {
	return &GroupController{
		log: log,
	}
}
