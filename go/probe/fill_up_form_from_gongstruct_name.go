// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongthree/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "BezierCurve":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "BezierCurve Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierCurveFormCallback(
			nil,
			probe,
			formGroup,
		)
		beziercurve := new(models.BezierCurve)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beziercurve, formGroup, probe)
	case "BezierSegment":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "BezierSegment Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierSegmentFormCallback(
			nil,
			probe,
			formGroup,
		)
		beziersegment := new(models.BezierSegment)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beziersegment, formGroup, probe)
	case "Vector2":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Vector2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Vector2FormCallback(
			nil,
			probe,
			formGroup,
		)
		vector2 := new(models.Vector2)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(vector2, formGroup, probe)
	}
	formStage.Commit()
}
