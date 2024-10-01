package transfer

import "github.com/google/uuid"

type StudentCreateRequest struct {
	FirstName        string
	LastName         string
	MiddleName       string
	Group            GroupData
	EducationalEmail string
	Username         string
	Password         string
}

type GroupData struct {
	SpecializationCode string
	GroupNumber        string
}

type StudentCreateResponse struct {
	StudentID uuid.UUID
}

type TeacherCreateRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	MiddleName  string `json:"middle_name"`
	ReportEmail string `json:"report_email" validate:"required,email"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type TeacherCreateResponse struct {
	TeacherID uuid.UUID
}
