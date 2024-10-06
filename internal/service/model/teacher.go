package model

import "github.com/google/uuid"

type Teacher struct {
	FirstName           string
	LastName            string
	MiddleName          string
	ReportEmail         string
	Username            string
	HashedPasswordBytes []byte
}

type TeacherVerify struct {
	TeacherID uuid.UUID
}
