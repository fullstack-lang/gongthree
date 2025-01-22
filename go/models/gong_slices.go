// generated code - do not edit
package models

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
