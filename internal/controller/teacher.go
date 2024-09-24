package controller

import transfer "github.com/utf2/utf-account-service/internal/controller/model"

type TeacherControllerAPI interface {
	Create(request transfer.TeacherCreateRequest) transfer.TeacherCreateResponse
	Verify(request transfer.TeacherVerifyRequest) transfer.TeacherVerifyResponse
	SearchByID(request transfer.TeacherSearchByIDRequest) transfer.TeacherSearchByIDResponse
}
