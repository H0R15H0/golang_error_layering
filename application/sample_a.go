package application

type SampleAError struct {
	// Commandで定義されたFieldに関連したエラー
	// 例: Field1が不正な値の場合はField1Errがセットされる
	Field1Err Error
	Field2Err Error
	// Commandで定義されたField以外に関連して発生するエラー
	// 例: Field1とField2の組み合わせが不正な場合はErrがセットされる
	// 例: DBに保存できない場合はErrがセットされる
	Err Error
}

// TODO: 適当に実装しただけ
func (e *SampleAError) Error() string {
	var retval string
	if e.Field1Err != nil {
		retval += e.Field1Err.Error()
	}
	if e.Field2Err != nil {
		retval += e.Field2Err.Error()
	}
	if e.Err != nil {
		retval += e.Err.Error()
	}
	return retval
}

type SampleACommand struct {
	Field1 string
	Field2 string
}

// func SampleA(cmd SampleACommand) error {
// 	field1, dErr := models.NewField1(cmd.Field1)
// 	field2, dErr := models.NewField2(cmd.Field2)

// 	sampleA, dErr := models.NewSampleA(field1, field2)

// 	dErr = repositories.NewSampleARepository().Save(sampleA)

// 	return nil
// }
