package transfer

import "github.com/google/uuid"

type StudentVerifyRequest struct {
	StudentID uuid.UUID
}

type TeacherVerifyRequest struct {
	TeacherID uuid.UUID `json:"teacher_id" validate:"required"`
}
