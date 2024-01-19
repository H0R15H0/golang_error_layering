package application_clean

import (
	"go_error/common"
	"go_error/domain/models"
	"go_error/domain/repositories"
)

type SaveSampleACommand struct {
	Field1 string
	Field2 string
}

type SaveSampleAOutputPort interface {
	SetField1Error(err Error)
	SetField2Error(err Error)
	SetError(err Error)
	Emit()
}

func SaveSampleA(cmd SaveSampleACommand, outputPort SaveSampleAOutputPort) (bool, error) {
	const op = "application_clean.SampleA"
	var hasError bool
	field1, dErr := models.NewField1(cmd.Field1)
	if dErr != nil {
		outputPort.SetField1Error(&DefaultError{
			err: &common.Error{
				Op:  op,
				Err: dErr,
			},
			userMessage: "フィールド1が不正です",
		})
		hasError = true
	}
	field2, dErr := models.NewField2(cmd.Field2)
	if dErr != nil {
		outputPort.SetField2Error(&DefaultError{
			err: &common.Error{
				Op:  op,
				Err: dErr,
			},
			userMessage: "フィールド2が不正です",
		})
		hasError = true
	}

	if hasError {
		return false, &common.Error{Op: op, Err: dErr}
	}

	sampleA, dErr := models.NewSampleA(field1, field2)
	if dErr != nil {
		outputPort.SetError(&DefaultError{
			err: &common.Error{
				Op:  op,
				Err: dErr,
			},
			userMessage: "サンプルAに一貫性がありません",
		})
		return false, &common.Error{Op: op, Err: dErr}
	}

	dErr = repositories.NewSampleARepository().Save(sampleA)
	if dErr != nil {
		outputPort.SetError(&DefaultError{
			err: &common.Error{
				Op:  op,
				Err: dErr,
			},
			userMessage: "予期せぬエラーが発生しました",
		})
		return false, &common.Error{Op: op, Err: dErr}
	}

	return true, nil
}
