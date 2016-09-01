package horizon

func (err *Error) Error() string {
	// TODO: use the attached problem to provide a better error message
	return "Horizon error"
}
