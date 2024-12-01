// generated code - do not edit

//insertion point for imports
import { BezierCurveAPI } from './beziercurve-api'

import { BezierSegmentAPI } from './beziersegment-api'

import { Vector2API } from './vector2-api'


export class BackRepoData {
	// insertion point for declarations
	BezierCurveAPIs = new Array<BezierCurveAPI>()

	BezierSegmentAPIs = new Array<BezierSegmentAPI>()

	Vector2APIs = new Array<Vector2API>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.BezierCurveAPIs = data?.BezierCurveAPIs || [];

		this.BezierSegmentAPIs = data?.BezierSegmentAPIs || [];

		this.Vector2APIs = data?.Vector2APIs || [];

	}

}