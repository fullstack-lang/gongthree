// generated code - do not edit

import { BezierCurveAPI } from './beziercurve-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { BezierSegment } from './beziersegment'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BezierCurve {

	static GONGSTRUCT_NAME = "BezierCurve"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Color: string = ""

	// insertion point for pointers and slices of pointers declarations
	BezierSegments: Array<BezierSegment> = []
}

export function CopyBezierCurveToBezierCurveAPI(beziercurve: BezierCurve, beziercurveAPI: BezierCurveAPI) {

	beziercurveAPI.CreatedAt = beziercurve.CreatedAt
	beziercurveAPI.DeletedAt = beziercurve.DeletedAt
	beziercurveAPI.ID = beziercurve.ID

	// insertion point for basic fields copy operations
	beziercurveAPI.Name = beziercurve.Name
	beziercurveAPI.Color = beziercurve.Color

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	beziercurveAPI.BezierCurvePointersEncoding.BezierSegments = []
	for (let _beziersegment of beziercurve.BezierSegments) {
		beziercurveAPI.BezierCurvePointersEncoding.BezierSegments.push(_beziersegment.ID)
	}

}

// CopyBezierCurveAPIToBezierCurve update basic, pointers and slice of pointers fields of beziercurve
// from respectively the basic fields and encoded fields of pointers and slices of pointers of beziercurveAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBezierCurveAPIToBezierCurve(beziercurveAPI: BezierCurveAPI, beziercurve: BezierCurve, frontRepo: FrontRepo) {

	beziercurve.CreatedAt = beziercurveAPI.CreatedAt
	beziercurve.DeletedAt = beziercurveAPI.DeletedAt
	beziercurve.ID = beziercurveAPI.ID

	// insertion point for basic fields copy operations
	beziercurve.Name = beziercurveAPI.Name
	beziercurve.Color = beziercurveAPI.Color

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(beziercurveAPI.BezierCurvePointersEncoding.BezierSegments)) {
		console.error('Rects is not an array:', beziercurveAPI.BezierCurvePointersEncoding.BezierSegments);
		return;
	}

	beziercurve.BezierSegments = new Array<BezierSegment>()
	for (let _id of beziercurveAPI.BezierCurvePointersEncoding.BezierSegments) {
		let _beziersegment = frontRepo.map_ID_BezierSegment.get(_id)
		if (_beziersegment != undefined) {
			beziercurve.BezierSegments.push(_beziersegment!)
		}
	}
}
