// generated code - do not edit

import { BezierSegmentAPI } from './beziersegment-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Vector2 } from './vector2'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BezierSegment {

	static GONGSTRUCT_NAME = "BezierSegment"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Start?: Vector2

	ControlPointStart?: Vector2

	ControlPointEnd?: Vector2

	End?: Vector2

}

export function CopyBezierSegmentToBezierSegmentAPI(beziersegment: BezierSegment, beziersegmentAPI: BezierSegmentAPI) {

	beziersegmentAPI.CreatedAt = beziersegment.CreatedAt
	beziersegmentAPI.DeletedAt = beziersegment.DeletedAt
	beziersegmentAPI.ID = beziersegment.ID

	// insertion point for basic fields copy operations
	beziersegmentAPI.Name = beziersegment.Name

	// insertion point for pointer fields encoding
	beziersegmentAPI.BezierSegmentPointersEncoding.StartID.Valid = true
	if (beziersegment.Start != undefined) {
		beziersegmentAPI.BezierSegmentPointersEncoding.StartID.Int64 = beziersegment.Start.ID  
	} else {
		beziersegmentAPI.BezierSegmentPointersEncoding.StartID.Int64 = 0 		
	}

	beziersegmentAPI.BezierSegmentPointersEncoding.ControlPointStartID.Valid = true
	if (beziersegment.ControlPointStart != undefined) {
		beziersegmentAPI.BezierSegmentPointersEncoding.ControlPointStartID.Int64 = beziersegment.ControlPointStart.ID  
	} else {
		beziersegmentAPI.BezierSegmentPointersEncoding.ControlPointStartID.Int64 = 0 		
	}

	beziersegmentAPI.BezierSegmentPointersEncoding.ControlPointEndID.Valid = true
	if (beziersegment.ControlPointEnd != undefined) {
		beziersegmentAPI.BezierSegmentPointersEncoding.ControlPointEndID.Int64 = beziersegment.ControlPointEnd.ID  
	} else {
		beziersegmentAPI.BezierSegmentPointersEncoding.ControlPointEndID.Int64 = 0 		
	}

	beziersegmentAPI.BezierSegmentPointersEncoding.EndID.Valid = true
	if (beziersegment.End != undefined) {
		beziersegmentAPI.BezierSegmentPointersEncoding.EndID.Int64 = beziersegment.End.ID  
	} else {
		beziersegmentAPI.BezierSegmentPointersEncoding.EndID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyBezierSegmentAPIToBezierSegment update basic, pointers and slice of pointers fields of beziersegment
// from respectively the basic fields and encoded fields of pointers and slices of pointers of beziersegmentAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBezierSegmentAPIToBezierSegment(beziersegmentAPI: BezierSegmentAPI, beziersegment: BezierSegment, frontRepo: FrontRepo) {

	beziersegment.CreatedAt = beziersegmentAPI.CreatedAt
	beziersegment.DeletedAt = beziersegmentAPI.DeletedAt
	beziersegment.ID = beziersegmentAPI.ID

	// insertion point for basic fields copy operations
	beziersegment.Name = beziersegmentAPI.Name

	// insertion point for pointer fields encoding
	beziersegment.Start = frontRepo.map_ID_Vector2.get(beziersegmentAPI.BezierSegmentPointersEncoding.StartID.Int64)
	beziersegment.ControlPointStart = frontRepo.map_ID_Vector2.get(beziersegmentAPI.BezierSegmentPointersEncoding.ControlPointStartID.Int64)
	beziersegment.ControlPointEnd = frontRepo.map_ID_Vector2.get(beziersegmentAPI.BezierSegmentPointersEncoding.ControlPointEndID.Int64)
	beziersegment.End = frontRepo.map_ID_Vector2.get(beziersegmentAPI.BezierSegmentPointersEncoding.EndID.Int64)

	// insertion point for slice of pointers fields encoding
}
