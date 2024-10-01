package controller

import (
	"log/slog"

	transfer "github.com/utf2/utf-account-service/internal/controller/model"
)

type StudentControllerAPI interface {
	Create(transfer.StudentCreateRequest) transfer.StudentCreateResponse
	Verify(transfer.StudentVerifyRequest)
	SearchByID(transfer.StudentSearchByIDRequest) transfer.StudentSearchByIDResponse
	SearchByGroupID(transfer.StudentSearchByGroupIDRequest) transfer.StudentSearchByGroupIDResponse
}

type StudentController struct {
	log *slog.Logger
}

func NewStudentController(log *slog.Logger) *StudentController {
	return &StudentController{
		log: log,
	}
}
