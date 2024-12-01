// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongthree/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.BezierCurve:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "BezierCurve Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierCurveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BezierSegment:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "BezierSegment Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierSegmentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Vector2:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Vector2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Vector2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
