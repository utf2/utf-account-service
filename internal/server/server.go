package server

import (
	"fmt"
	"log/slog"
	"net/http"

	chi "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	config "github.com/utf2/utf-account-service/internal/config/app"
	"github.com/utf2/utf-account-service/internal/controller"
	"github.com/utf2/utf-account-service/internal/handler"
	"github.com/utf2/utf-account-service/internal/middleware"

	_ "github.com/utf2/utf-account-service/docs"
)

func New(config *config.Config, log *slog.Logger) (*http.Server, error) {
	teacherController := controller.NewTeacherController(log, nil)
	_ = controller.NewStudentController(log)
	_ = controller.NewGroupController(log)

	userMux := http.NewServeMux()
	userMux.HandleFunc("POST /teacher", handler.WrapHttpHandlerFunc(teacherController.Create))
	userMux.HandleFunc("POST /teacher/verify", handler.WrapHttpHandlerFunc(teacherController.Create))
	userMux.HandleFunc("POST /teacher/{id}", handler.WrapHttpHandlerFunc(teacherController.Create))

	userMux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	apiV1 := http.NewServeMux()
	apiV1.Handle("/api/v1/", http.StripPrefix("/api/v1", userMux))

	stack := middleware.CreateStack(
		chi.RequestID,
		middleware.Logging(log),
		chi.Recoverer,
	)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", config.HttpServer.Port),
		Handler: stack(apiV1),
	}, nil
}
