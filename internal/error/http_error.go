package httperror

type HttpError struct {
	Err          error
	StatusCode   int
	ErrorMessage string
}

func (err *HttpError) Error() string {
	return err.Err.Error()
}
