// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongthree/go/models"
	"github.com/fullstack-lang/gongthree/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.BezierCurve:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("BezierSegments", instanceWithInferedType, &instanceWithInferedType.BezierSegments, formGroup, probe)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.BezierSegment:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, probe)
		AssociationFieldToForm("ControlPointStart", instanceWithInferedType.ControlPointStart, formGroup, probe)
		AssociationFieldToForm("ControlPointEnd", instanceWithInferedType.ControlPointEnd, formGroup, probe)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "BezierCurve"
			rf.Fieldname = "BezierSegments"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.BezierCurve),
					"BezierSegments",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.BezierCurve, *models.BezierSegment](
					nil,
					"BezierSegments",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Vector2:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	default:
		_ = instanceWithInferedType
	}
}
