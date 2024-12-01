package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongthree/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__Field__000000_X := (&models.Field{}).Stage(stage)
	__Field__000001_Y := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_BezierCurve := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_BezierSegment := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Vector2 := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_BezierSegments := (&models.Link{}).Stage(stage)
	__Link__000001_ControlPointEnd := (&models.Link{}).Stage(stage)
	__Link__000002_ControlPointStart := (&models.Link{}).Stage(stage)
	__Link__000003_End := (&models.Link{}).Stage(stage)
	__Link__000004_Start := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_BezierCurve := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_BezierSegment := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_Vector2 := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierCurve_and_Default_BezierSegment := (&models.Vertice{}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2 := (&models.Vertice{}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2 := (&models.Vertice{}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2 := (&models.Vertice{}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2 := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_X.Name = `X`

	//gong:ident [ref_models.Vector2.X] comment added to overcome the problem with the comment map association
	__Field__000000_X.Identifier = `ref_models.Vector2.X`
	__Field__000000_X.FieldTypeAsString = ``
	__Field__000000_X.Structname = `Vector2`
	__Field__000000_X.Fieldtypename = `float64`

	__Field__000001_Y.Name = `Y`

	//gong:ident [ref_models.Vector2.Y] comment added to overcome the problem with the comment map association
	__Field__000001_Y.Identifier = `ref_models.Vector2.Y`
	__Field__000001_Y.FieldTypeAsString = ``
	__Field__000001_Y.Structname = `Vector2`
	__Field__000001_Y.Fieldtypename = `float64`

	__GongStructShape__000000_Default_BezierCurve.Name = `Default-BezierCurve`

	//gong:ident [ref_models.BezierCurve] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_BezierCurve.Identifier = `ref_models.BezierCurve`
	__GongStructShape__000000_Default_BezierCurve.ShowNbInstances = false
	__GongStructShape__000000_Default_BezierCurve.NbInstances = 0
	__GongStructShape__000000_Default_BezierCurve.Width = 240.000000
	__GongStructShape__000000_Default_BezierCurve.Height = 63.000000
	__GongStructShape__000000_Default_BezierCurve.IsSelected = false

	__GongStructShape__000001_Default_BezierSegment.Name = `Default-BezierSegment`

	//gong:ident [ref_models.BezierSegment] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_BezierSegment.Identifier = `ref_models.BezierSegment`
	__GongStructShape__000001_Default_BezierSegment.ShowNbInstances = false
	__GongStructShape__000001_Default_BezierSegment.NbInstances = 0
	__GongStructShape__000001_Default_BezierSegment.Width = 240.000000
	__GongStructShape__000001_Default_BezierSegment.Height = 63.000000
	__GongStructShape__000001_Default_BezierSegment.IsSelected = false

	__GongStructShape__000002_Default_Vector2.Name = `Default-Vector2`

	//gong:ident [ref_models.Vector2] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Vector2.Identifier = `ref_models.Vector2`
	__GongStructShape__000002_Default_Vector2.ShowNbInstances = false
	__GongStructShape__000002_Default_Vector2.NbInstances = 0
	__GongStructShape__000002_Default_Vector2.Width = 240.000000
	__GongStructShape__000002_Default_Vector2.Height = 305.000000
	__GongStructShape__000002_Default_Vector2.IsSelected = false

	__Link__000000_BezierSegments.Name = `BezierSegments`

	//gong:ident [ref_models.BezierCurve.BezierSegments] comment added to overcome the problem with the comment map association
	__Link__000000_BezierSegments.Identifier = `ref_models.BezierCurve.BezierSegments`

	//gong:ident [ref_models.BezierSegment] comment added to overcome the problem with the comment map association
	__Link__000000_BezierSegments.Fieldtypename = `ref_models.BezierSegment`
	__Link__000000_BezierSegments.FieldOffsetX = -50.000000
	__Link__000000_BezierSegments.FieldOffsetY = -16.000000
	__Link__000000_BezierSegments.TargetMultiplicity = models.MANY
	__Link__000000_BezierSegments.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_BezierSegments.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_BezierSegments.SourceMultiplicity = models.MANY
	__Link__000000_BezierSegments.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_BezierSegments.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_BezierSegments.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_BezierSegments.StartRatio = 0.500000
	__Link__000000_BezierSegments.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_BezierSegments.EndRatio = 0.500000
	__Link__000000_BezierSegments.CornerOffsetRatio = 1.380000

	__Link__000001_ControlPointEnd.Name = `ControlPointEnd`

	//gong:ident [ref_models.BezierSegment.ControlPointEnd] comment added to overcome the problem with the comment map association
	__Link__000001_ControlPointEnd.Identifier = `ref_models.BezierSegment.ControlPointEnd`

	//gong:ident [ref_models.Vector2] comment added to overcome the problem with the comment map association
	__Link__000001_ControlPointEnd.Fieldtypename = `ref_models.Vector2`
	__Link__000001_ControlPointEnd.FieldOffsetX = -50.000000
	__Link__000001_ControlPointEnd.FieldOffsetY = -16.000000
	__Link__000001_ControlPointEnd.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_ControlPointEnd.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_ControlPointEnd.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_ControlPointEnd.SourceMultiplicity = models.MANY
	__Link__000001_ControlPointEnd.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_ControlPointEnd.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_ControlPointEnd.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_ControlPointEnd.StartRatio = 0.500000
	__Link__000001_ControlPointEnd.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_ControlPointEnd.EndRatio = 0.674753
	__Link__000001_ControlPointEnd.CornerOffsetRatio = 1.380000

	__Link__000002_ControlPointStart.Name = `ControlPointStart`

	//gong:ident [ref_models.BezierSegment.ControlPointStart] comment added to overcome the problem with the comment map association
	__Link__000002_ControlPointStart.Identifier = `ref_models.BezierSegment.ControlPointStart`

	//gong:ident [ref_models.Vector2] comment added to overcome the problem with the comment map association
	__Link__000002_ControlPointStart.Fieldtypename = `ref_models.Vector2`
	__Link__000002_ControlPointStart.FieldOffsetX = -50.000000
	__Link__000002_ControlPointStart.FieldOffsetY = -16.000000
	__Link__000002_ControlPointStart.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_ControlPointStart.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_ControlPointStart.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_ControlPointStart.SourceMultiplicity = models.MANY
	__Link__000002_ControlPointStart.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_ControlPointStart.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_ControlPointStart.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_ControlPointStart.StartRatio = 0.500000
	__Link__000002_ControlPointStart.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_ControlPointStart.EndRatio = 0.333770
	__Link__000002_ControlPointStart.CornerOffsetRatio = 1.380000

	__Link__000003_End.Name = `End`

	//gong:ident [ref_models.BezierSegment.End] comment added to overcome the problem with the comment map association
	__Link__000003_End.Identifier = `ref_models.BezierSegment.End`

	//gong:ident [ref_models.Vector2] comment added to overcome the problem with the comment map association
	__Link__000003_End.Fieldtypename = `ref_models.Vector2`
	__Link__000003_End.FieldOffsetX = -50.000000
	__Link__000003_End.FieldOffsetY = -16.000000
	__Link__000003_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_End.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_End.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_End.SourceMultiplicity = models.MANY
	__Link__000003_End.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_End.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_End.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_End.StartRatio = 0.500000
	__Link__000003_End.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_End.EndRatio = 0.973114
	__Link__000003_End.CornerOffsetRatio = 1.380000

	__Link__000004_Start.Name = `Start`

	//gong:ident [ref_models.BezierSegment.Start] comment added to overcome the problem with the comment map association
	__Link__000004_Start.Identifier = `ref_models.BezierSegment.Start`

	//gong:ident [ref_models.Vector2] comment added to overcome the problem with the comment map association
	__Link__000004_Start.Fieldtypename = `ref_models.Vector2`
	__Link__000004_Start.FieldOffsetX = -50.000000
	__Link__000004_Start.FieldOffsetY = -16.000000
	__Link__000004_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_Start.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_Start.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_Start.SourceMultiplicity = models.MANY
	__Link__000004_Start.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_Start.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_Start.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_Start.StartRatio = 0.500000
	__Link__000004_Start.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_Start.EndRatio = 0.009180
	__Link__000004_Start.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_BezierCurve.X = 56.000000
	__Position__000000_Pos_Default_BezierCurve.Y = 75.000000
	__Position__000000_Pos_Default_BezierCurve.Name = `Pos-Default-BezierCurve`

	__Position__000001_Pos_Default_BezierSegment.X = 500.000000
	__Position__000001_Pos_Default_BezierSegment.Y = 73.000000
	__Position__000001_Pos_Default_BezierSegment.Name = `Pos-Default-BezierSegment`

	__Position__000002_Pos_Default_Vector2.X = 1033.999969
	__Position__000002_Pos_Default_Vector2.Y = 93.000000
	__Position__000002_Pos_Default_Vector2.Name = `Pos-Default-Vector2`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierCurve_and_Default_BezierSegment.X = 578.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierCurve_and_Default_BezierSegment.Y = 110.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierCurve_and_Default_BezierSegment.Name = `Verticle in class diagram Default in middle between Default-BezierCurve and Default-BezierSegment`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.X = 1080.499985
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.Y = 103.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.Name = `Verticle in class diagram Default in middle between Default-BezierSegment and Default-Vector2`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.X = 1080.499985
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.Y = 103.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.Name = `Verticle in class diagram Default in middle between Default-BezierSegment and Default-Vector2`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.X = 1079.999985
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.Y = 103.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.Name = `Verticle in class diagram Default in middle between Default-BezierSegment and Default-Vector2`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.X = 1079.999985
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.Y = 103.500000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2.Name = `Verticle in class diagram Default in middle between Default-BezierSegment and Default-Vector2`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_BezierCurve)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_BezierSegment)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Vector2)
	__GongStructShape__000000_Default_BezierCurve.Position = __Position__000000_Pos_Default_BezierCurve
	__GongStructShape__000000_Default_BezierCurve.Links = append(__GongStructShape__000000_Default_BezierCurve.Links, __Link__000000_BezierSegments)
	__GongStructShape__000001_Default_BezierSegment.Position = __Position__000001_Pos_Default_BezierSegment
	__GongStructShape__000001_Default_BezierSegment.Links = append(__GongStructShape__000001_Default_BezierSegment.Links, __Link__000004_Start)
	__GongStructShape__000001_Default_BezierSegment.Links = append(__GongStructShape__000001_Default_BezierSegment.Links, __Link__000002_ControlPointStart)
	__GongStructShape__000001_Default_BezierSegment.Links = append(__GongStructShape__000001_Default_BezierSegment.Links, __Link__000001_ControlPointEnd)
	__GongStructShape__000001_Default_BezierSegment.Links = append(__GongStructShape__000001_Default_BezierSegment.Links, __Link__000003_End)
	__GongStructShape__000002_Default_Vector2.Position = __Position__000002_Pos_Default_Vector2
	__GongStructShape__000002_Default_Vector2.Fields = append(__GongStructShape__000002_Default_Vector2.Fields, __Field__000000_X)
	__GongStructShape__000002_Default_Vector2.Fields = append(__GongStructShape__000002_Default_Vector2.Fields, __Field__000001_Y)
	__Link__000000_BezierSegments.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierCurve_and_Default_BezierSegment
	__Link__000001_ControlPointEnd.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2
	__Link__000002_ControlPointStart.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2
	__Link__000003_End.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2
	__Link__000004_Start.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_BezierSegment_and_Default_Vector2
}
