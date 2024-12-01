// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongthree/go/models"
	"github.com/fullstack-lang/gongthree/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__CountryFormCallback(
	country *models.Country,
	probe *Probe,
	formGroup *table.FormGroup,
) (countryFormCallback *CountryFormCallback) {
	countryFormCallback = new(CountryFormCallback)
	countryFormCallback.probe = probe
	countryFormCallback.country = country
	countryFormCallback.formGroup = formGroup

	countryFormCallback.CreationMode = (country == nil)

	return
}

type CountryFormCallback struct {
	country *models.Country

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (countryFormCallback *CountryFormCallback) OnSave() {

	log.Println("CountryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	countryFormCallback.probe.formStage.Checkout()

	if countryFormCallback.country == nil {
		countryFormCallback.country = new(models.Country).Stage(countryFormCallback.probe.stageOfInterest)
	}
	country_ := countryFormCallback.country
	_ = country_

	for _, formDiv := range countryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(country_.Name), formDiv)
		case "Hello":
			FormDivSelectFieldToField(&(country_.Hello), countryFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if countryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		country_.Unstage(countryFormCallback.probe.stageOfInterest)
	}

	countryFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Country](
		countryFormCallback.probe,
	)
	countryFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if countryFormCallback.CreationMode || countryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		countryFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(countryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CountryFormCallback(
			nil,
			countryFormCallback.probe,
			newFormGroup,
		)
		country := new(models.Country)
		FillUpForm(country, newFormGroup, countryFormCallback.probe)
		countryFormCallback.probe.formStage.Commit()
	}

	fillUpTree(countryFormCallback.probe)
}
func __gong__New__HelloFormCallback(
	hello *models.Hello,
	probe *Probe,
	formGroup *table.FormGroup,
) (helloFormCallback *HelloFormCallback) {
	helloFormCallback = new(HelloFormCallback)
	helloFormCallback.probe = probe
	helloFormCallback.hello = hello
	helloFormCallback.formGroup = formGroup

	helloFormCallback.CreationMode = (hello == nil)

	return
}

type HelloFormCallback struct {
	hello *models.Hello

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (helloFormCallback *HelloFormCallback) OnSave() {

	log.Println("HelloFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	helloFormCallback.probe.formStage.Checkout()

	if helloFormCallback.hello == nil {
		helloFormCallback.hello = new(models.Hello).Stage(helloFormCallback.probe.stageOfInterest)
	}
	hello_ := helloFormCallback.hello
	_ = hello_

	for _, formDiv := range helloFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(hello_.Name), formDiv)
		case "Country:AlternateHellos":
			// we need to retrieve the field owner before the change
			var pastCountryOwner *models.Country
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Country"
			rf.Fieldname = "AlternateHellos"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				helloFormCallback.probe.stageOfInterest,
				helloFormCallback.probe.backRepoOfInterest,
				hello_,
				&rf)

			if reverseFieldOwner != nil {
				pastCountryOwner = reverseFieldOwner.(*models.Country)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastCountryOwner != nil {
					idx := slices.Index(pastCountryOwner.AlternateHellos, hello_)
					pastCountryOwner.AlternateHellos = slices.Delete(pastCountryOwner.AlternateHellos, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _country := range *models.GetGongstructInstancesSet[models.Country](helloFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _country.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newCountryOwner := _country // we have a match
						if pastCountryOwner != nil {
							if newCountryOwner != pastCountryOwner {
								idx := slices.Index(pastCountryOwner.AlternateHellos, hello_)
								pastCountryOwner.AlternateHellos = slices.Delete(pastCountryOwner.AlternateHellos, idx, idx+1)
								newCountryOwner.AlternateHellos = append(newCountryOwner.AlternateHellos, hello_)
							}
						} else {
							newCountryOwner.AlternateHellos = append(newCountryOwner.AlternateHellos, hello_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if helloFormCallback.formGroup.HasSuppressButtonBeenPressed {
		hello_.Unstage(helloFormCallback.probe.stageOfInterest)
	}

	helloFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Hello](
		helloFormCallback.probe,
	)
	helloFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if helloFormCallback.CreationMode || helloFormCallback.formGroup.HasSuppressButtonBeenPressed {
		helloFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(helloFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__HelloFormCallback(
			nil,
			helloFormCallback.probe,
			newFormGroup,
		)
		hello := new(models.Hello)
		FillUpForm(hello, newFormGroup, helloFormCallback.probe)
		helloFormCallback.probe.formStage.Commit()
	}

	fillUpTree(helloFormCallback.probe)
}
