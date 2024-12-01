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
func __gong__New__BezierCurveFormCallback(
	beziercurve *models.BezierCurve,
	probe *Probe,
	formGroup *table.FormGroup,
) (beziercurveFormCallback *BezierCurveFormCallback) {
	beziercurveFormCallback = new(BezierCurveFormCallback)
	beziercurveFormCallback.probe = probe
	beziercurveFormCallback.beziercurve = beziercurve
	beziercurveFormCallback.formGroup = formGroup

	beziercurveFormCallback.CreationMode = (beziercurve == nil)

	return
}

type BezierCurveFormCallback struct {
	beziercurve *models.BezierCurve

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (beziercurveFormCallback *BezierCurveFormCallback) OnSave() {

	log.Println("BezierCurveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	beziercurveFormCallback.probe.formStage.Checkout()

	if beziercurveFormCallback.beziercurve == nil {
		beziercurveFormCallback.beziercurve = new(models.BezierCurve).Stage(beziercurveFormCallback.probe.stageOfInterest)
	}
	beziercurve_ := beziercurveFormCallback.beziercurve
	_ = beziercurve_

	for _, formDiv := range beziercurveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(beziercurve_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if beziercurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziercurve_.Unstage(beziercurveFormCallback.probe.stageOfInterest)
	}

	beziercurveFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.BezierCurve](
		beziercurveFormCallback.probe,
	)
	beziercurveFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beziercurveFormCallback.CreationMode || beziercurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziercurveFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(beziercurveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BezierCurveFormCallback(
			nil,
			beziercurveFormCallback.probe,
			newFormGroup,
		)
		beziercurve := new(models.BezierCurve)
		FillUpForm(beziercurve, newFormGroup, beziercurveFormCallback.probe)
		beziercurveFormCallback.probe.formStage.Commit()
	}

	fillUpTree(beziercurveFormCallback.probe)
}
func __gong__New__BezierSegmentFormCallback(
	beziersegment *models.BezierSegment,
	probe *Probe,
	formGroup *table.FormGroup,
) (beziersegmentFormCallback *BezierSegmentFormCallback) {
	beziersegmentFormCallback = new(BezierSegmentFormCallback)
	beziersegmentFormCallback.probe = probe
	beziersegmentFormCallback.beziersegment = beziersegment
	beziersegmentFormCallback.formGroup = formGroup

	beziersegmentFormCallback.CreationMode = (beziersegment == nil)

	return
}

type BezierSegmentFormCallback struct {
	beziersegment *models.BezierSegment

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (beziersegmentFormCallback *BezierSegmentFormCallback) OnSave() {

	log.Println("BezierSegmentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	beziersegmentFormCallback.probe.formStage.Checkout()

	if beziersegmentFormCallback.beziersegment == nil {
		beziersegmentFormCallback.beziersegment = new(models.BezierSegment).Stage(beziersegmentFormCallback.probe.stageOfInterest)
	}
	beziersegment_ := beziersegmentFormCallback.beziersegment
	_ = beziersegment_

	for _, formDiv := range beziersegmentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(beziersegment_.Name), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(beziersegment_.Start), beziersegmentFormCallback.probe.stageOfInterest, formDiv)
		case "ControlPointStart":
			FormDivSelectFieldToField(&(beziersegment_.ControlPointStart), beziersegmentFormCallback.probe.stageOfInterest, formDiv)
		case "ControlPointEnd":
			FormDivSelectFieldToField(&(beziersegment_.ControlPointEnd), beziersegmentFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(beziersegment_.End), beziersegmentFormCallback.probe.stageOfInterest, formDiv)
		case "BezierCurve:BezierSegments":
			// we need to retrieve the field owner before the change
			var pastBezierCurveOwner *models.BezierCurve
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "BezierCurve"
			rf.Fieldname = "BezierSegments"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				beziersegmentFormCallback.probe.stageOfInterest,
				beziersegmentFormCallback.probe.backRepoOfInterest,
				beziersegment_,
				&rf)

			if reverseFieldOwner != nil {
				pastBezierCurveOwner = reverseFieldOwner.(*models.BezierCurve)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastBezierCurveOwner != nil {
					idx := slices.Index(pastBezierCurveOwner.BezierSegments, beziersegment_)
					pastBezierCurveOwner.BezierSegments = slices.Delete(pastBezierCurveOwner.BezierSegments, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _beziercurve := range *models.GetGongstructInstancesSet[models.BezierCurve](beziersegmentFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _beziercurve.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newBezierCurveOwner := _beziercurve // we have a match
						if pastBezierCurveOwner != nil {
							if newBezierCurveOwner != pastBezierCurveOwner {
								idx := slices.Index(pastBezierCurveOwner.BezierSegments, beziersegment_)
								pastBezierCurveOwner.BezierSegments = slices.Delete(pastBezierCurveOwner.BezierSegments, idx, idx+1)
								newBezierCurveOwner.BezierSegments = append(newBezierCurveOwner.BezierSegments, beziersegment_)
							}
						} else {
							newBezierCurveOwner.BezierSegments = append(newBezierCurveOwner.BezierSegments, beziersegment_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if beziersegmentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziersegment_.Unstage(beziersegmentFormCallback.probe.stageOfInterest)
	}

	beziersegmentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.BezierSegment](
		beziersegmentFormCallback.probe,
	)
	beziersegmentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beziersegmentFormCallback.CreationMode || beziersegmentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziersegmentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(beziersegmentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BezierSegmentFormCallback(
			nil,
			beziersegmentFormCallback.probe,
			newFormGroup,
		)
		beziersegment := new(models.BezierSegment)
		FillUpForm(beziersegment, newFormGroup, beziersegmentFormCallback.probe)
		beziersegmentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(beziersegmentFormCallback.probe)
}
func __gong__New__Vector2FormCallback(
	vector2 *models.Vector2,
	probe *Probe,
	formGroup *table.FormGroup,
) (vector2FormCallback *Vector2FormCallback) {
	vector2FormCallback = new(Vector2FormCallback)
	vector2FormCallback.probe = probe
	vector2FormCallback.vector2 = vector2
	vector2FormCallback.formGroup = formGroup

	vector2FormCallback.CreationMode = (vector2 == nil)

	return
}

type Vector2FormCallback struct {
	vector2 *models.Vector2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (vector2FormCallback *Vector2FormCallback) OnSave() {

	log.Println("Vector2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	vector2FormCallback.probe.formStage.Checkout()

	if vector2FormCallback.vector2 == nil {
		vector2FormCallback.vector2 = new(models.Vector2).Stage(vector2FormCallback.probe.stageOfInterest)
	}
	vector2_ := vector2FormCallback.vector2
	_ = vector2_

	for _, formDiv := range vector2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(vector2_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(vector2_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(vector2_.Y), formDiv)
		}
	}

	// manage the suppress operation
	if vector2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		vector2_.Unstage(vector2FormCallback.probe.stageOfInterest)
	}

	vector2FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Vector2](
		vector2FormCallback.probe,
	)
	vector2FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if vector2FormCallback.CreationMode || vector2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		vector2FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(vector2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Vector2FormCallback(
			nil,
			vector2FormCallback.probe,
			newFormGroup,
		)
		vector2 := new(models.Vector2)
		FillUpForm(vector2, newFormGroup, vector2FormCallback.probe)
		vector2FormCallback.probe.formStage.Commit()
	}

	fillUpTree(vector2FormCallback.probe)
}
