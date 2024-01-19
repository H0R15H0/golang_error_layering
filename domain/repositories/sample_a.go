package repositories

import (
	"fmt"
	"go_error/common"
	"go_error/domain"
	"go_error/domain/models"
)

type SampleARepository interface {
	Save(sampleA models.SampleA) domain.Error
}

type sampleARepository struct {
}

func NewSampleARepository() SampleARepository {
	return &sampleARepository{}
}

func (r *sampleARepository) Save(sampleA models.SampleA) domain.Error {
	op := "repositories.sampleARepository.Save"
	if sampleA.ShouldNotSave {
		return domain.NewDefaultError(&common.Error{
			Op:      op,
			Code:    common.ErrorCodeInvalidArgument,
			Message: "SampleA is invalid",
		})
	}
	fmt.Println("SampleA is saved!")
	return nil
}
