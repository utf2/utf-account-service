package model

type Teacher struct {
	FirstName           string
	LastName            string
	MiddleName          string
	ReportEmail         string
	Username            string
	HashedPasswordBytes []byte
}
