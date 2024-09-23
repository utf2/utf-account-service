package transfer

import "github.com/google/uuid"

type GroupDTO struct {
	SpecializationCode string
	GroupNumber        string
}

type StudentDTO struct {
	ID               uuid.UUID
	FirstName        string
	LastName         string
	MiddleName       string
	EducationalEmail string
	Username         string
	Group            GroupDTO
}

type StudentSearchByIDRequest struct {
	StudentID uuid.UUID
}

type StudentSearchByIDResponse struct {
	StudentData StudentDTO
}

type TeacherDTO struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	MiddleName   string
	ReportEmail  string
	ContactEmail string
	Username     string
}

type TeacherSearchByIDRequest struct {
	TeacherID uuid.UUID
}

type TeacherSearchByIDResponse struct {
	TeacherData TeacherDTO
}
