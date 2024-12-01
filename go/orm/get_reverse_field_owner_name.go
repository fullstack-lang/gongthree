// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongthree/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Country:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Hello:
		switch reverseField.GongstructName {
		// insertion point
		case "Country":
			switch reverseField.Fieldname {
			case "AlternateHellos":
				if _country, ok := stage.Country_AlternateHellos_reverseMap[inst]; ok {
					res = _country.Name
				}
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Country:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Hello:
		switch reverseField.GongstructName {
		// insertion point
		case "Country":
			switch reverseField.Fieldname {
			case "AlternateHellos":
				res = stage.Country_AlternateHellos_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
