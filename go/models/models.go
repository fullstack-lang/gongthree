package models

type Vector2 struct {
	Name string
	X, Y float64
}

type BezierSegment struct {
	Name              string
	Start             *Vector2
	ControlPointStart *Vector2
	ControlPointEnd   *Vector2
	End               *Vector2
}

type BezierCurve struct {
	Name string

	BezierSegments []*BezierSegment

	Color string
}
