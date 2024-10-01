package converter

import (
	transfer "github.com/utf2/utf-account-service/internal/controller/model"
	"github.com/utf2/utf-account-service/internal/service/model"
)

func ConvertTeacherCreateRequest(request transfer.TeacherCreateRequest) model.Teacher {
	return model.Teacher{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		MiddleName:  request.MiddleName,
		ReportEmail: request.ReportEmail,
		Username:    request.Username,
	}
}
