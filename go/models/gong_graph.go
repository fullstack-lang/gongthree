// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *BezierCurve:
		ok = stage.IsStagedBezierCurve(target)

	case *BezierSegment:
		ok = stage.IsStagedBezierSegment(target)

	case *Vector2:
		ok = stage.IsStagedVector2(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedBezierCurve(beziercurve *BezierCurve) (ok bool) {

	_, ok = stage.BezierCurves[beziercurve]

	return
}

func (stage *StageStruct) IsStagedBezierSegment(beziersegment *BezierSegment) (ok bool) {

	_, ok = stage.BezierSegments[beziersegment]

	return
}

func (stage *StageStruct) IsStagedVector2(vector2 *Vector2) (ok bool) {

	_, ok = stage.Vector2s[vector2]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *BezierCurve:
		stage.StageBranchBezierCurve(target)

	case *BezierSegment:
		stage.StageBranchBezierSegment(target)

	case *Vector2:
		stage.StageBranchVector2(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchBezierCurve(beziercurve *BezierCurve) {

	// check if instance is already staged
	if IsStaged(stage, beziercurve) {
		return
	}

	beziercurve.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beziersegment := range beziercurve.BezierSegments {
		StageBranch(stage, _beziersegment)
	}

}

func (stage *StageStruct) StageBranchBezierSegment(beziersegment *BezierSegment) {

	// check if instance is already staged
	if IsStaged(stage, beziersegment) {
		return
	}

	beziersegment.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if beziersegment.Start != nil {
		StageBranch(stage, beziersegment.Start)
	}
	if beziersegment.ControlPointStart != nil {
		StageBranch(stage, beziersegment.ControlPointStart)
	}
	if beziersegment.ControlPointEnd != nil {
		StageBranch(stage, beziersegment.ControlPointEnd)
	}
	if beziersegment.End != nil {
		StageBranch(stage, beziersegment.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchVector2(vector2 *Vector2) {

	// check if instance is already staged
	if IsStaged(stage, vector2) {
		return
	}

	vector2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *BezierCurve:
		toT := CopyBranchBezierCurve(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BezierSegment:
		toT := CopyBranchBezierSegment(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Vector2:
		toT := CopyBranchVector2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchBezierCurve(mapOrigCopy map[any]any, beziercurveFrom *BezierCurve) (beziercurveTo *BezierCurve) {

	// beziercurveFrom has already been copied
	if _beziercurveTo, ok := mapOrigCopy[beziercurveFrom]; ok {
		beziercurveTo = _beziercurveTo.(*BezierCurve)
		return
	}

	beziercurveTo = new(BezierCurve)
	mapOrigCopy[beziercurveFrom] = beziercurveTo
	beziercurveFrom.CopyBasicFields(beziercurveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beziersegment := range beziercurveFrom.BezierSegments {
		beziercurveTo.BezierSegments = append(beziercurveTo.BezierSegments, CopyBranchBezierSegment(mapOrigCopy, _beziersegment))
	}

	return
}

func CopyBranchBezierSegment(mapOrigCopy map[any]any, beziersegmentFrom *BezierSegment) (beziersegmentTo *BezierSegment) {

	// beziersegmentFrom has already been copied
	if _beziersegmentTo, ok := mapOrigCopy[beziersegmentFrom]; ok {
		beziersegmentTo = _beziersegmentTo.(*BezierSegment)
		return
	}

	beziersegmentTo = new(BezierSegment)
	mapOrigCopy[beziersegmentFrom] = beziersegmentTo
	beziersegmentFrom.CopyBasicFields(beziersegmentTo)

	//insertion point for the staging of instances referenced by pointers
	if beziersegmentFrom.Start != nil {
		beziersegmentTo.Start = CopyBranchVector2(mapOrigCopy, beziersegmentFrom.Start)
	}
	if beziersegmentFrom.ControlPointStart != nil {
		beziersegmentTo.ControlPointStart = CopyBranchVector2(mapOrigCopy, beziersegmentFrom.ControlPointStart)
	}
	if beziersegmentFrom.ControlPointEnd != nil {
		beziersegmentTo.ControlPointEnd = CopyBranchVector2(mapOrigCopy, beziersegmentFrom.ControlPointEnd)
	}
	if beziersegmentFrom.End != nil {
		beziersegmentTo.End = CopyBranchVector2(mapOrigCopy, beziersegmentFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVector2(mapOrigCopy map[any]any, vector2From *Vector2) (vector2To *Vector2) {

	// vector2From has already been copied
	if _vector2To, ok := mapOrigCopy[vector2From]; ok {
		vector2To = _vector2To.(*Vector2)
		return
	}

	vector2To = new(Vector2)
	mapOrigCopy[vector2From] = vector2To
	vector2From.CopyBasicFields(vector2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *BezierCurve:
		stage.UnstageBranchBezierCurve(target)

	case *BezierSegment:
		stage.UnstageBranchBezierSegment(target)

	case *Vector2:
		stage.UnstageBranchVector2(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchBezierCurve(beziercurve *BezierCurve) {

	// check if instance is already staged
	if !IsStaged(stage, beziercurve) {
		return
	}

	beziercurve.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beziersegment := range beziercurve.BezierSegments {
		UnstageBranch(stage, _beziersegment)
	}

}

func (stage *StageStruct) UnstageBranchBezierSegment(beziersegment *BezierSegment) {

	// check if instance is already staged
	if !IsStaged(stage, beziersegment) {
		return
	}

	beziersegment.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if beziersegment.Start != nil {
		UnstageBranch(stage, beziersegment.Start)
	}
	if beziersegment.ControlPointStart != nil {
		UnstageBranch(stage, beziersegment.ControlPointStart)
	}
	if beziersegment.ControlPointEnd != nil {
		UnstageBranch(stage, beziersegment.ControlPointEnd)
	}
	if beziersegment.End != nil {
		UnstageBranch(stage, beziersegment.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchVector2(vector2 *Vector2) {

	// check if instance is already staged
	if !IsStaged(stage, vector2) {
		return
	}

	vector2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
