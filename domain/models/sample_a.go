package models

import (
	"go_error/common"
	"go_error/domain"
)

type Field1 struct {
	value string
}

var ZeroField1 = Field1{}

func NewField1(value string) (Field1, domain.Error) {
	const op = "domain.NewField1"
	if value == "invalidField1" {
		return ZeroField1, domain.NewDefaultError(&common.Error{
			Op:      op,
			Code:    common.ErrorCodeInvalidArgument,
			Message: "Field1 is invalid",
		})
	}
	return Field1{value: value}, nil
}

type Field2 struct {
	value string
}

var ZeroField2 = Field2{}

func NewField2(value string) (Field2, domain.Error) {
	const op = "domain.NewField2"
	if value == "invalidField2" {
		return ZeroField2, domain.NewDefaultError(&common.Error{
			Op:      op,
			Code:    common.ErrorCodeInvalidArgument,
			Message: "Field2 is invalid",
		})
	}
	return Field2{value: value}, nil
}

type SampleA struct {
	field1        Field1
	field2        Field2
	ShouldNotSave bool
}

var ZeroSampleA = SampleA{
	field1:        ZeroField1,
	field2:        ZeroField2,
	ShouldNotSave: true,
}

func NewSampleA(field1 Field1, field2 Field2) (SampleA, domain.Error) {
	const op = "domain.NewSampleA"
	if field1.value == "entityError" || field2.value == "entityError" {
		return ZeroSampleA, domain.NewDefaultError(&common.Error{
			Op:      op,
			Code:    common.ErrorCodeInvalidArgument,
			Message: "SampleA is invalid",
		})
	}
	return SampleA{field1: field1, field2: field2}, nil
}
