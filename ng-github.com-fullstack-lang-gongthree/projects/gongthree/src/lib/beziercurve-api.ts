// insertion point for imports
import { BezierSegmentAPI } from './beziersegment-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BezierCurveAPI {

	static GONGSTRUCT_NAME = "BezierCurve"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Color: string = ""

	// insertion point for other decls

	BezierCurvePointersEncoding: BezierCurvePointersEncoding = new BezierCurvePointersEncoding
}

export class BezierCurvePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	BezierSegments: number[] = []
}
