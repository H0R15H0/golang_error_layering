package graphclean

type UserError struct {
	Fields  []string
	Message string
}

type SaveSampleAResponse struct {
	UserErrors []UserError
}
