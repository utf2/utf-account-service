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
	FirstName   string
	LastName    string
	MiddleName  string
	ReportEmail string
	Username    string
	Password    string
}

type TeacherCreateResponse struct {
	TeacherID uuid.UUID
}
