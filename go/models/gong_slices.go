// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *BezierCurve:
		// insertion point per field
		if fieldName == "BezierSegments" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*BezierCurve) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*BezierCurve)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.BezierSegments).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.BezierSegments = _inferedTypeInstance.BezierSegments[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.BezierSegments =
								append(_inferedTypeInstance.BezierSegments, any(fieldInstance).(*BezierSegment))
						}
					}
				}
			}
		}

	case *BezierSegment:
		// insertion point per field

	case *Vector2:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct BezierCurve
	// insertion point per field
	clear(stage.BezierCurve_BezierSegments_reverseMap)
	stage.BezierCurve_BezierSegments_reverseMap = make(map[*BezierSegment]*BezierCurve)
	for beziercurve := range stage.BezierCurves {
		_ = beziercurve
		for _, _beziersegment := range beziercurve.BezierSegments {
			stage.BezierCurve_BezierSegments_reverseMap[_beziersegment] = beziercurve
		}
	}

	// Compute reverse map for named struct BezierSegment
	// insertion point per field

	// Compute reverse map for named struct Vector2
	// insertion point per field

}
