package pkg

import (
	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/pkg/errs"
	"github.com/asaskevich/govalidator"
)

func ValidateStruct(payload interface{}) errs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}
	return nil
}
