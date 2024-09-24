package controller

import transfer "github.com/utf2/utf-account-service/internal/controller/model"

type StudentControllerAPI interface {
	Create(transfer.StudentCreateRequest) transfer.StudentCreateResponse
	Verify(transfer.StudentVerifyRequest) transfer.StudentVerifyResponse
	SearchByID(transfer.StudentSearchByIDRequest) transfer.StudentSearchByIDResponse
	SearchByGroupID(transfer.StudentSearchByGroupIDRequest) transfer.StudentSearchByGroupIDResponse
}
