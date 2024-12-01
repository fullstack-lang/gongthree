// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type BezierCurve_WOP struct {
	// insertion point
	Name string
}

func (from *BezierCurve) CopyBasicFields(to *BezierCurve) {
	// insertion point
	to.Name = from.Name
}

type BezierSegment_WOP struct {
	// insertion point
	Name string
}

func (from *BezierSegment) CopyBasicFields(to *BezierSegment) {
	// insertion point
	to.Name = from.Name
}

type Vector2_WOP struct {
	// insertion point
	Name string
	X float64
	Y float64
}

func (from *Vector2) CopyBasicFields(to *Vector2) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
}

