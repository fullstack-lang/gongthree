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
	case "Country":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Country Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CountryFormCallback(
			nil,
			probe,
			formGroup,
		)
		country := new(models.Country)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(country, formGroup, probe)
	case "Hello":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Hello Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HelloFormCallback(
			nil,
			probe,
			formGroup,
		)
		hello := new(models.Hello)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(hello, formGroup, probe)
	}
	formStage.Commit()
}
