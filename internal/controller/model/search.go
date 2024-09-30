package transfer

import "github.com/google/uuid"

type StudentSearchByIDRequest struct {
	StudentID uuid.UUID
}

type StudentSearchByIDResponse struct {
	StudentData StudentDTO
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

type GroupDTO struct {
	ID                 uuid.UUID
	SpecializationCode string
	GroupNumber        string
}

type StudentSearchByGroupIDRequest struct {
	GroupID uuid.UUID
}

type StudentSearchByGroupIDResponse struct {
	StudentsInGroup []StudentDTO
}

type TeacherSearchByIDRequest struct {
	TeacherID uuid.UUID
}

type TeacherSearchByIDResponse struct {
	TeacherData TeacherDTO
}

type TeacherDTO struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	MiddleName  string
	ReportEmail string
	Username    string
}

type GroupSearchRequest struct {
	SpecializationCode string
	GroupNumber        string
}

type GroupSearchResponse struct {
	MatchedGroups []GroupDTO
}
