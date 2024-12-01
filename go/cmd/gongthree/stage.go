package main

import (
	"time"

	"github.com/fullstack-lang/gongthree/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__BezierCurve__000000_Test := (&models.BezierCurve{}).Stage(stage)

	__BezierSegment__000000_S1 := (&models.BezierSegment{}).Stage(stage)
	__BezierSegment__000001_S2 := (&models.BezierSegment{}).Stage(stage)

	__Vector2__000000_S1_E := (&models.Vector2{}).Stage(stage)
	__Vector2__000001_S1_E_CP := (&models.Vector2{}).Stage(stage)
	__Vector2__000002_S1_S := (&models.Vector2{}).Stage(stage)
	__Vector2__000003_S1_S_CP := (&models.Vector2{}).Stage(stage)
	__Vector2__000004_S2_E := (&models.Vector2{}).Stage(stage)
	__Vector2__000005_S2_E_CP := (&models.Vector2{}).Stage(stage)
	__Vector2__000006_S2_S := (&models.Vector2{}).Stage(stage)
	__Vector2__000007_S2_S_CP := (&models.Vector2{}).Stage(stage)

	// Setup of values

	__BezierCurve__000000_Test.Name = `Test`
	__BezierCurve__000000_Test.Color = `#8FA382`

	__BezierSegment__000000_S1.Name = `S1`

	__BezierSegment__000001_S2.Name = `S2`

	__Vector2__000000_S1_E.Name = `S1 E`
	__Vector2__000000_S1_E.X = -5.000000
	__Vector2__000000_S1_E.Y = 4.000000

	__Vector2__000001_S1_E_CP.Name = `S1 E CP`
	__Vector2__000001_S1_E_CP.X = -2.000000
	__Vector2__000001_S1_E_CP.Y = 8.000000

	__Vector2__000002_S1_S.Name = `S1 S`
	__Vector2__000002_S1_S.X = 5.000000
	__Vector2__000002_S1_S.Y = 4.000000

	__Vector2__000003_S1_S_CP.Name = `S1 S CP`
	__Vector2__000003_S1_S_CP.X = 2.000000
	__Vector2__000003_S1_S_CP.Y = 8.000000

	__Vector2__000004_S2_E.Name = `S2 E`
	__Vector2__000004_S2_E.X = -10.000000
	__Vector2__000004_S2_E.Y = 0.000000

	__Vector2__000005_S2_E_CP.Name = `S2 E CP`
	__Vector2__000005_S2_E_CP.X = -7.000000
	__Vector2__000005_S2_E_CP.Y = 4.000000

	__Vector2__000006_S2_S.Name = `S2 S`
	__Vector2__000006_S2_S.X = -5.000000
	__Vector2__000006_S2_S.Y = 4.000000

	__Vector2__000007_S2_S_CP.Name = `S2 S CP`
	__Vector2__000007_S2_S_CP.X = -8.000000
	__Vector2__000007_S2_S_CP.Y = 0.000000

	// Setup of pointers
	__BezierCurve__000000_Test.BezierSegments = append(__BezierCurve__000000_Test.BezierSegments, __BezierSegment__000000_S1)
	__BezierCurve__000000_Test.BezierSegments = append(__BezierCurve__000000_Test.BezierSegments, __BezierSegment__000001_S2)
	__BezierSegment__000000_S1.Start = __Vector2__000002_S1_S
	__BezierSegment__000000_S1.ControlPointStart = __Vector2__000003_S1_S_CP
	__BezierSegment__000000_S1.ControlPointEnd = __Vector2__000001_S1_E_CP
	__BezierSegment__000000_S1.End = __Vector2__000000_S1_E
	__BezierSegment__000001_S2.Start = __Vector2__000006_S2_S
	__BezierSegment__000001_S2.ControlPointStart = __Vector2__000007_S2_S_CP
	__BezierSegment__000001_S2.ControlPointEnd = __Vector2__000005_S2_E_CP
	__BezierSegment__000001_S2.End = __Vector2__000004_S2_E
}
