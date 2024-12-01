// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongthree/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.BezierCurve:
		beziercurveInstance := any(concreteInstance).(*models.BezierCurve)
		ret2 := backRepo.BackRepoBezierCurve.GetBezierCurveDBFromBezierCurvePtr(beziercurveInstance)
		ret = any(ret2).(*T2)
	case *models.BezierSegment:
		beziersegmentInstance := any(concreteInstance).(*models.BezierSegment)
		ret2 := backRepo.BackRepoBezierSegment.GetBezierSegmentDBFromBezierSegmentPtr(beziersegmentInstance)
		ret = any(ret2).(*T2)
	case *models.Vector2:
		vector2Instance := any(concreteInstance).(*models.Vector2)
		ret2 := backRepo.BackRepoVector2.GetVector2DBFromVector2Ptr(vector2Instance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.BezierCurve:
		tmp := GetInstanceDBFromInstance[models.BezierCurve, BezierCurveDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BezierSegment:
		tmp := GetInstanceDBFromInstance[models.BezierSegment, BezierSegmentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Vector2:
		tmp := GetInstanceDBFromInstance[models.Vector2, Vector2DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.BezierCurve:
		tmp := GetInstanceDBFromInstance[models.BezierCurve, BezierCurveDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BezierSegment:
		tmp := GetInstanceDBFromInstance[models.BezierSegment, BezierSegmentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Vector2:
		tmp := GetInstanceDBFromInstance[models.Vector2, Vector2DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
