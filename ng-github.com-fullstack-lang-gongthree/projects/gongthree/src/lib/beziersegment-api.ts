// insertion point for imports
import { Vector2API } from './vector2-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BezierSegmentAPI {

	static GONGSTRUCT_NAME = "BezierSegment"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	BezierSegmentPointersEncoding: BezierSegmentPointersEncoding = new BezierSegmentPointersEncoding
}

export class BezierSegmentPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	StartID: NullInt64 = new NullInt64 // if pointer is null, Start.ID = 0

	ControlPointStartID: NullInt64 = new NullInt64 // if pointer is null, ControlPointStart.ID = 0

	ControlPointEndID: NullInt64 = new NullInt64 // if pointer is null, ControlPointEnd.ID = 0

	EndID: NullInt64 = new NullInt64 // if pointer is null, End.ID = 0

}
