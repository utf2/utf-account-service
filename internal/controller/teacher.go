package controller

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/utf2/utf-account-service/internal/controller/converter"
	transfer "github.com/utf2/utf-account-service/internal/controller/model"
	httperror "github.com/utf2/utf-account-service/internal/error"
	"github.com/utf2/utf-account-service/internal/logger"
	"github.com/utf2/utf-account-service/internal/service/model"
)

type TeacherController struct {
	log     *slog.Logger
	service TeacherService
}

func NewTeacherController(log *slog.Logger, service TeacherService) *TeacherController {
	return &TeacherController{
		log:     log,
		service: service,
	}
}

type TeacherService interface {
	Save(context.Context, model.Teacher) (uuid.UUID, error)
	Verify(context.Context, model.TeacherVerify) error
}

// @Summary create a new teacher
// @Accept json
// @Produce json
// @Param teacherData body transfer.TeacherCreateRequest true "Required field values to create teacher"
// @Success 200 {object} transfer.TeacherCreateResponse
// @Router /api/v1/teacher [post]
func (controller *TeacherController) Create(w http.ResponseWriter, r *http.Request) *httperror.HttpError {
	const op = "controller.TeacherController.Create()"

	log := controller.log.With(
		slog.String("op", op),
		slog.String("requestID", middleware.GetReqID(r.Context())),
	)

	var request transfer.TeacherCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error("error while decoding teacher create request body from JSON", logger.Error(err))
		return &httperror.HttpError{
			Err:          err,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "Error while decoding teacher create request body from JSON",
		}
	}

	if err := validator.New().Struct(request); err != nil {
		log.Error("error while validating teacher create request body from JSON", logger.Error(err))
		return &httperror.HttpError{
			Err:          err,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "Error while validating teacher create request body from JSON",
		}
	}

	teacher := converter.ConvertTeacherCreateRequest(request)
	ctx := context.WithValue(context.Background(), middleware.RequestIDKey, middleware.GetReqID(r.Context()))
	createdTeacherId, err := controller.service.Save(ctx, teacher)
	if err != nil {
		return &httperror.HttpError{
			Err:          err,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "Error while saving teacher data",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transfer.TeacherCreateResponse{
		TeacherID: createdTeacherId,
	})

	return nil
}

// @Summary verify teacher account
// @Accept json
// @Produce json
// @Param teacherID body transfer.TeacherVerifyRequest true "Request body with verifying teacher ID"
// @Success 200
// @Router /api/v1/teacher/verify [post]
func (controller *TeacherController) Verify(w http.ResponseWriter, r *http.Request) *httperror.HttpError {
	const op = "controller.TeacherController.Verify()"

	log := controller.log.With(
		slog.String("op", op),
		slog.String("requestID", middleware.GetReqID(r.Context())),
	)

	var request transfer.TeacherVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error("error while decoding teacher verify request body from JSON", logger.Error(err))
		return &httperror.HttpError{
			Err:          err,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "Error while decoding teacher verify request body from JSON",
		}
	}

	if err := validator.New().Struct(request); err != nil {
		log.Error("error while validating teacher verify request body from JSON", logger.Error(err))
		return &httperror.HttpError{
			Err:          err,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "Error while validating teacher verify request body from JSON",
		}
	}

	teacherVerify := converter.ConvertTeacherVerifyRequest(request)
	ctx := context.WithValue(context.Background(), middleware.RequestIDKey, middleware.GetReqID(r.Context()))
	if err := controller.service.Verify(ctx, teacherVerify); err != nil {
		return &httperror.HttpError{
			Err:          err,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "Error while verifying teacher",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}

func (controller *TeacherController) SearchByID(w http.ResponseWriter, r *http.Request) *httperror.HttpError {
	panic("not implemented!")
}
