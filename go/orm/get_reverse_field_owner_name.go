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
	case *models.BezierCurve:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.BezierSegment:
		switch reverseField.GongstructName {
		// insertion point
		case "BezierCurve":
			switch reverseField.Fieldname {
			case "BezierSegments":
				if _beziercurve, ok := stage.BezierCurve_BezierSegments_reverseMap[inst]; ok {
					res = _beziercurve.Name
				}
			}
		}

	case *models.Vector2:
		switch reverseField.GongstructName {
		// insertion point
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
	case *models.BezierCurve:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.BezierSegment:
		switch reverseField.GongstructName {
		// insertion point
		case "BezierCurve":
			switch reverseField.Fieldname {
			case "BezierSegments":
				res = stage.BezierCurve_BezierSegments_reverseMap[inst]
			}
		}

	case *models.Vector2:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
