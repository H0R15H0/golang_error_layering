package graphclean

import "go_error/application_clean"

type SaveSampleAOutPort struct {
	Field1Error application_clean.Error
	Field2Error application_clean.Error
	otherError  application_clean.Error
	Response    *SaveSampleAResponse
}

func (o *SaveSampleAOutPort) SetField1Error(err application_clean.Error) {
	o.Field1Error = err
}

func (o *SaveSampleAOutPort) SetField2Error(err application_clean.Error) {
	o.Field2Error = err
}

func (o *SaveSampleAOutPort) SetError(err application_clean.Error) {
	o.otherError = err
}

func (o *SaveSampleAOutPort) Emit() {
	o.Response = &SaveSampleAResponse{
		UserErrors: []UserError{},
	}
	if o.Field1Error != nil {
		o.Response.UserErrors = append(o.Response.UserErrors, UserError{
			Fields:  []string{"sampleA", "field1"},
			Message: o.Field1Error.UserMessage(),
		})
	}
	if o.Field2Error != nil {
		o.Response.UserErrors = append(o.Response.UserErrors, UserError{
			Fields:  []string{"sampleA", "field2"},
			Message: o.Field2Error.UserMessage(),
		})
	}
	if o.otherError != nil {
		o.Response.UserErrors = append(o.Response.UserErrors, UserError{
			Fields:  []string{},
			Message: o.otherError.UserMessage(),
		})
	}
}
