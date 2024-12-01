// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	BezierCurveAPIs []*BezierCurveAPI

	BezierSegmentAPIs []*BezierSegmentAPI

	Vector2APIs []*Vector2API
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, beziercurveDB := range backRepo.BackRepoBezierCurve.Map_BezierCurveDBID_BezierCurveDB {

		var beziercurveAPI BezierCurveAPI
		beziercurveAPI.ID = beziercurveDB.ID
		beziercurveAPI.BezierCurvePointersEncoding = beziercurveDB.BezierCurvePointersEncoding
		beziercurveDB.CopyBasicFieldsToBezierCurve_WOP(&beziercurveAPI.BezierCurve_WOP)

		backRepoData.BezierCurveAPIs = append(backRepoData.BezierCurveAPIs, &beziercurveAPI)
	}

	for _, beziersegmentDB := range backRepo.BackRepoBezierSegment.Map_BezierSegmentDBID_BezierSegmentDB {

		var beziersegmentAPI BezierSegmentAPI
		beziersegmentAPI.ID = beziersegmentDB.ID
		beziersegmentAPI.BezierSegmentPointersEncoding = beziersegmentDB.BezierSegmentPointersEncoding
		beziersegmentDB.CopyBasicFieldsToBezierSegment_WOP(&beziersegmentAPI.BezierSegment_WOP)

		backRepoData.BezierSegmentAPIs = append(backRepoData.BezierSegmentAPIs, &beziersegmentAPI)
	}

	for _, vector2DB := range backRepo.BackRepoVector2.Map_Vector2DBID_Vector2DB {

		var vector2API Vector2API
		vector2API.ID = vector2DB.ID
		vector2API.Vector2PointersEncoding = vector2DB.Vector2PointersEncoding
		vector2DB.CopyBasicFieldsToVector2_WOP(&vector2API.Vector2_WOP)

		backRepoData.Vector2APIs = append(backRepoData.Vector2APIs, &vector2API)
	}

}
