package transfer

import "github.com/google/uuid"

type StudentVerifyRequest struct {
	StudentID uuid.UUID
}

type StudentVerifyResponse struct {
	Success bool
}

type TeacherVerifyRequest struct {
	TeacherID uuid.UUID
}

type TeacherVerifyResponse struct {
	Success bool
}
