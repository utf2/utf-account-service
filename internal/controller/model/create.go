package transfer

import "github.com/google/uuid"

type GroupCreateRequest struct {
	SpecializationCode string
	GroupNumber        string
}

type GroupCreateResponse struct {
	Success bool
	GroupID uuid.UUID
}

type StudentCreateRequest struct {
	FirstName        string
	LastName         string
	MiddleName       string
	EducationalEmail string
	Username         string
	Password         string
	GroupID          uuid.UUID
}

type StudentRegisterResponse struct {
	Success   bool
	StudentID uuid.UUID
}

type TeacherRegisterRequest struct {
	FirstName    string
	LastName     string
	MiddleName   string
	ReportEmail  string
	ContactEmail string
	Username     string
	Password     string
}

type TeacherRegisterResponse struct {
	Success   bool
	TeacherID uuid.UUID
}
