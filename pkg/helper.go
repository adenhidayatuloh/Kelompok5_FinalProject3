package pkg

import (
	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/pkg/errs"
	"github.com/asaskevich/govalidator"
)

func ValidateStruct(payload interface{}) errs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}
	return nil
}
