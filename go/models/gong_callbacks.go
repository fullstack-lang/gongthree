// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *BezierCurve:
		if stage.OnAfterBezierCurveCreateCallback != nil {
			stage.OnAfterBezierCurveCreateCallback.OnAfterCreate(stage, target)
		}
	case *BezierSegment:
		if stage.OnAfterBezierSegmentCreateCallback != nil {
			stage.OnAfterBezierSegmentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Vector2:
		if stage.OnAfterVector2CreateCallback != nil {
			stage.OnAfterVector2CreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *BezierCurve:
		newTarget := any(new).(*BezierCurve)
		if stage.OnAfterBezierCurveUpdateCallback != nil {
			stage.OnAfterBezierCurveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BezierSegment:
		newTarget := any(new).(*BezierSegment)
		if stage.OnAfterBezierSegmentUpdateCallback != nil {
			stage.OnAfterBezierSegmentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Vector2:
		newTarget := any(new).(*Vector2)
		if stage.OnAfterVector2UpdateCallback != nil {
			stage.OnAfterVector2UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *BezierCurve:
		if stage.OnAfterBezierCurveDeleteCallback != nil {
			staged := any(staged).(*BezierCurve)
			stage.OnAfterBezierCurveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BezierSegment:
		if stage.OnAfterBezierSegmentDeleteCallback != nil {
			staged := any(staged).(*BezierSegment)
			stage.OnAfterBezierSegmentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Vector2:
		if stage.OnAfterVector2DeleteCallback != nil {
			staged := any(staged).(*Vector2)
			stage.OnAfterVector2DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *BezierCurve:
		if stage.OnAfterBezierCurveReadCallback != nil {
			stage.OnAfterBezierCurveReadCallback.OnAfterRead(stage, target)
		}
	case *BezierSegment:
		if stage.OnAfterBezierSegmentReadCallback != nil {
			stage.OnAfterBezierSegmentReadCallback.OnAfterRead(stage, target)
		}
	case *Vector2:
		if stage.OnAfterVector2ReadCallback != nil {
			stage.OnAfterVector2ReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BezierCurve:
		stage.OnAfterBezierCurveUpdateCallback = any(callback).(OnAfterUpdateInterface[BezierCurve])
	
	case *BezierSegment:
		stage.OnAfterBezierSegmentUpdateCallback = any(callback).(OnAfterUpdateInterface[BezierSegment])
	
	case *Vector2:
		stage.OnAfterVector2UpdateCallback = any(callback).(OnAfterUpdateInterface[Vector2])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BezierCurve:
		stage.OnAfterBezierCurveCreateCallback = any(callback).(OnAfterCreateInterface[BezierCurve])
	
	case *BezierSegment:
		stage.OnAfterBezierSegmentCreateCallback = any(callback).(OnAfterCreateInterface[BezierSegment])
	
	case *Vector2:
		stage.OnAfterVector2CreateCallback = any(callback).(OnAfterCreateInterface[Vector2])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BezierCurve:
		stage.OnAfterBezierCurveDeleteCallback = any(callback).(OnAfterDeleteInterface[BezierCurve])
	
	case *BezierSegment:
		stage.OnAfterBezierSegmentDeleteCallback = any(callback).(OnAfterDeleteInterface[BezierSegment])
	
	case *Vector2:
		stage.OnAfterVector2DeleteCallback = any(callback).(OnAfterDeleteInterface[Vector2])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BezierCurve:
		stage.OnAfterBezierCurveReadCallback = any(callback).(OnAfterReadInterface[BezierCurve])
	
	case *BezierSegment:
		stage.OnAfterBezierSegmentReadCallback = any(callback).(OnAfterReadInterface[BezierSegment])
	
	case *Vector2:
		stage.OnAfterVector2ReadCallback = any(callback).(OnAfterReadInterface[Vector2])
	
	}
}
