// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	BezierCurves           map[*BezierCurve]any
	BezierCurves_mapString map[string]*BezierCurve

	// insertion point for slice of pointers maps
	BezierCurve_BezierSegments_reverseMap map[*BezierSegment]*BezierCurve

	OnAfterBezierCurveCreateCallback OnAfterCreateInterface[BezierCurve]
	OnAfterBezierCurveUpdateCallback OnAfterUpdateInterface[BezierCurve]
	OnAfterBezierCurveDeleteCallback OnAfterDeleteInterface[BezierCurve]
	OnAfterBezierCurveReadCallback   OnAfterReadInterface[BezierCurve]

	BezierSegments           map[*BezierSegment]any
	BezierSegments_mapString map[string]*BezierSegment

	// insertion point for slice of pointers maps
	OnAfterBezierSegmentCreateCallback OnAfterCreateInterface[BezierSegment]
	OnAfterBezierSegmentUpdateCallback OnAfterUpdateInterface[BezierSegment]
	OnAfterBezierSegmentDeleteCallback OnAfterDeleteInterface[BezierSegment]
	OnAfterBezierSegmentReadCallback   OnAfterReadInterface[BezierSegment]

	Vector2s           map[*Vector2]any
	Vector2s_mapString map[string]*Vector2

	// insertion point for slice of pointers maps
	OnAfterVector2CreateCallback OnAfterCreateInterface[Vector2]
	OnAfterVector2UpdateCallback OnAfterUpdateInterface[Vector2]
	OnAfterVector2DeleteCallback OnAfterDeleteInterface[Vector2]
	OnAfterVector2ReadCallback   OnAfterReadInterface[Vector2]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongthree/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitBezierCurve(beziercurve *BezierCurve)
	CheckoutBezierCurve(beziercurve *BezierCurve)
	CommitBezierSegment(beziersegment *BezierSegment)
	CheckoutBezierSegment(beziersegment *BezierSegment)
	CommitVector2(vector2 *Vector2)
	CheckoutVector2(vector2 *Vector2)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		BezierCurves:           make(map[*BezierCurve]any),
		BezierCurves_mapString: make(map[string]*BezierCurve),

		BezierSegments:           make(map[*BezierSegment]any),
		BezierSegments_mapString: make(map[string]*BezierSegment),

		Vector2s:           make(map[*Vector2]any),
		Vector2s_mapString: make(map[string]*Vector2),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["BezierCurve"] = len(stage.BezierCurves)
	stage.Map_GongStructName_InstancesNb["BezierSegment"] = len(stage.BezierSegments)
	stage.Map_GongStructName_InstancesNb["Vector2"] = len(stage.Vector2s)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["BezierCurve"] = len(stage.BezierCurves)
	stage.Map_GongStructName_InstancesNb["BezierSegment"] = len(stage.BezierSegments)
	stage.Map_GongStructName_InstancesNb["Vector2"] = len(stage.Vector2s)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts beziercurve to the model stage
func (beziercurve *BezierCurve) Stage(stage *StageStruct) *BezierCurve {
	stage.BezierCurves[beziercurve] = __member
	stage.BezierCurves_mapString[beziercurve.Name] = beziercurve

	return beziercurve
}

// Unstage removes beziercurve off the model stage
func (beziercurve *BezierCurve) Unstage(stage *StageStruct) *BezierCurve {
	delete(stage.BezierCurves, beziercurve)
	delete(stage.BezierCurves_mapString, beziercurve.Name)
	return beziercurve
}

// UnstageVoid removes beziercurve off the model stage
func (beziercurve *BezierCurve) UnstageVoid(stage *StageStruct) {
	delete(stage.BezierCurves, beziercurve)
	delete(stage.BezierCurves_mapString, beziercurve.Name)
}

// commit beziercurve to the back repo (if it is already staged)
func (beziercurve *BezierCurve) Commit(stage *StageStruct) *BezierCurve {
	if _, ok := stage.BezierCurves[beziercurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBezierCurve(beziercurve)
		}
	}
	return beziercurve
}

func (beziercurve *BezierCurve) CommitVoid(stage *StageStruct) {
	beziercurve.Commit(stage)
}

// Checkout beziercurve to the back repo (if it is already staged)
func (beziercurve *BezierCurve) Checkout(stage *StageStruct) *BezierCurve {
	if _, ok := stage.BezierCurves[beziercurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBezierCurve(beziercurve)
		}
	}
	return beziercurve
}

// for satisfaction of GongStruct interface
func (beziercurve *BezierCurve) GetName() (res string) {
	return beziercurve.Name
}

// Stage puts beziersegment to the model stage
func (beziersegment *BezierSegment) Stage(stage *StageStruct) *BezierSegment {
	stage.BezierSegments[beziersegment] = __member
	stage.BezierSegments_mapString[beziersegment.Name] = beziersegment

	return beziersegment
}

// Unstage removes beziersegment off the model stage
func (beziersegment *BezierSegment) Unstage(stage *StageStruct) *BezierSegment {
	delete(stage.BezierSegments, beziersegment)
	delete(stage.BezierSegments_mapString, beziersegment.Name)
	return beziersegment
}

// UnstageVoid removes beziersegment off the model stage
func (beziersegment *BezierSegment) UnstageVoid(stage *StageStruct) {
	delete(stage.BezierSegments, beziersegment)
	delete(stage.BezierSegments_mapString, beziersegment.Name)
}

// commit beziersegment to the back repo (if it is already staged)
func (beziersegment *BezierSegment) Commit(stage *StageStruct) *BezierSegment {
	if _, ok := stage.BezierSegments[beziersegment]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBezierSegment(beziersegment)
		}
	}
	return beziersegment
}

func (beziersegment *BezierSegment) CommitVoid(stage *StageStruct) {
	beziersegment.Commit(stage)
}

// Checkout beziersegment to the back repo (if it is already staged)
func (beziersegment *BezierSegment) Checkout(stage *StageStruct) *BezierSegment {
	if _, ok := stage.BezierSegments[beziersegment]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBezierSegment(beziersegment)
		}
	}
	return beziersegment
}

// for satisfaction of GongStruct interface
func (beziersegment *BezierSegment) GetName() (res string) {
	return beziersegment.Name
}

// Stage puts vector2 to the model stage
func (vector2 *Vector2) Stage(stage *StageStruct) *Vector2 {
	stage.Vector2s[vector2] = __member
	stage.Vector2s_mapString[vector2.Name] = vector2

	return vector2
}

// Unstage removes vector2 off the model stage
func (vector2 *Vector2) Unstage(stage *StageStruct) *Vector2 {
	delete(stage.Vector2s, vector2)
	delete(stage.Vector2s_mapString, vector2.Name)
	return vector2
}

// UnstageVoid removes vector2 off the model stage
func (vector2 *Vector2) UnstageVoid(stage *StageStruct) {
	delete(stage.Vector2s, vector2)
	delete(stage.Vector2s_mapString, vector2.Name)
}

// commit vector2 to the back repo (if it is already staged)
func (vector2 *Vector2) Commit(stage *StageStruct) *Vector2 {
	if _, ok := stage.Vector2s[vector2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitVector2(vector2)
		}
	}
	return vector2
}

func (vector2 *Vector2) CommitVoid(stage *StageStruct) {
	vector2.Commit(stage)
}

// Checkout vector2 to the back repo (if it is already staged)
func (vector2 *Vector2) Checkout(stage *StageStruct) *Vector2 {
	if _, ok := stage.Vector2s[vector2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutVector2(vector2)
		}
	}
	return vector2
}

// for satisfaction of GongStruct interface
func (vector2 *Vector2) GetName() (res string) {
	return vector2.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMBezierCurve(BezierCurve *BezierCurve)
	CreateORMBezierSegment(BezierSegment *BezierSegment)
	CreateORMVector2(Vector2 *Vector2)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMBezierCurve(BezierCurve *BezierCurve)
	DeleteORMBezierSegment(BezierSegment *BezierSegment)
	DeleteORMVector2(Vector2 *Vector2)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.BezierCurves = make(map[*BezierCurve]any)
	stage.BezierCurves_mapString = make(map[string]*BezierCurve)

	stage.BezierSegments = make(map[*BezierSegment]any)
	stage.BezierSegments_mapString = make(map[string]*BezierSegment)

	stage.Vector2s = make(map[*Vector2]any)
	stage.Vector2s_mapString = make(map[string]*Vector2)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.BezierCurves = nil
	stage.BezierCurves_mapString = nil

	stage.BezierSegments = nil
	stage.BezierSegments_mapString = nil

	stage.Vector2s = nil
	stage.Vector2s_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for beziercurve := range stage.BezierCurves {
		beziercurve.Unstage(stage)
	}

	for beziersegment := range stage.BezierSegments {
		beziersegment.Unstage(stage)
	}

	for vector2 := range stage.Vector2s {
		vector2.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*BezierCurve]any:
		return any(&stage.BezierCurves).(*Type)
	case map[*BezierSegment]any:
		return any(&stage.BezierSegments).(*Type)
	case map[*Vector2]any:
		return any(&stage.Vector2s).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*BezierCurve:
		return any(&stage.BezierCurves_mapString).(*Type)
	case map[string]*BezierSegment:
		return any(&stage.BezierSegments_mapString).(*Type)
	case map[string]*Vector2:
		return any(&stage.Vector2s_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case BezierCurve:
		return any(&stage.BezierCurves).(*map[*Type]any)
	case BezierSegment:
		return any(&stage.BezierSegments).(*map[*Type]any)
	case Vector2:
		return any(&stage.Vector2s).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *BezierCurve:
		return any(&stage.BezierCurves).(*map[Type]any)
	case *BezierSegment:
		return any(&stage.BezierSegments).(*map[Type]any)
	case *Vector2:
		return any(&stage.Vector2s).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case BezierCurve:
		return any(&stage.BezierCurves_mapString).(*map[string]*Type)
	case BezierSegment:
		return any(&stage.BezierSegments_mapString).(*map[string]*Type)
	case Vector2:
		return any(&stage.Vector2s_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case BezierCurve:
		return any(&BezierCurve{
			// Initialisation of associations
			// field is initialized with an instance of BezierSegment with the name of the field
			BezierSegments: []*BezierSegment{{Name: "BezierSegments"}},
		}).(*Type)
	case BezierSegment:
		return any(&BezierSegment{
			// Initialisation of associations
			// field is initialized with an instance of Vector2 with the name of the field
			Start: &Vector2{Name: "Start"},
			// field is initialized with an instance of Vector2 with the name of the field
			ControlPointStart: &Vector2{Name: "ControlPointStart"},
			// field is initialized with an instance of Vector2 with the name of the field
			ControlPointEnd: &Vector2{Name: "ControlPointEnd"},
			// field is initialized with an instance of Vector2 with the name of the field
			End: &Vector2{Name: "End"},
		}).(*Type)
	case Vector2:
		return any(&Vector2{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of BezierCurve
	case BezierCurve:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BezierSegment
	case BezierSegment:
		switch fieldname {
		// insertion point for per direct association field
		case "Start":
			res := make(map[*Vector2][]*BezierSegment)
			for beziersegment := range stage.BezierSegments {
				if beziersegment.Start != nil {
					vector2_ := beziersegment.Start
					var beziersegments []*BezierSegment
					_, ok := res[vector2_]
					if ok {
						beziersegments = res[vector2_]
					} else {
						beziersegments = make([]*BezierSegment, 0)
					}
					beziersegments = append(beziersegments, beziersegment)
					res[vector2_] = beziersegments
				}
			}
			return any(res).(map[*End][]*Start)
		case "ControlPointStart":
			res := make(map[*Vector2][]*BezierSegment)
			for beziersegment := range stage.BezierSegments {
				if beziersegment.ControlPointStart != nil {
					vector2_ := beziersegment.ControlPointStart
					var beziersegments []*BezierSegment
					_, ok := res[vector2_]
					if ok {
						beziersegments = res[vector2_]
					} else {
						beziersegments = make([]*BezierSegment, 0)
					}
					beziersegments = append(beziersegments, beziersegment)
					res[vector2_] = beziersegments
				}
			}
			return any(res).(map[*End][]*Start)
		case "ControlPointEnd":
			res := make(map[*Vector2][]*BezierSegment)
			for beziersegment := range stage.BezierSegments {
				if beziersegment.ControlPointEnd != nil {
					vector2_ := beziersegment.ControlPointEnd
					var beziersegments []*BezierSegment
					_, ok := res[vector2_]
					if ok {
						beziersegments = res[vector2_]
					} else {
						beziersegments = make([]*BezierSegment, 0)
					}
					beziersegments = append(beziersegments, beziersegment)
					res[vector2_] = beziersegments
				}
			}
			return any(res).(map[*End][]*Start)
		case "End":
			res := make(map[*Vector2][]*BezierSegment)
			for beziersegment := range stage.BezierSegments {
				if beziersegment.End != nil {
					vector2_ := beziersegment.End
					var beziersegments []*BezierSegment
					_, ok := res[vector2_]
					if ok {
						beziersegments = res[vector2_]
					} else {
						beziersegments = make([]*BezierSegment, 0)
					}
					beziersegments = append(beziersegments, beziersegment)
					res[vector2_] = beziersegments
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Vector2
	case Vector2:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of BezierCurve
	case BezierCurve:
		switch fieldname {
		// insertion point for per direct association field
		case "BezierSegments":
			res := make(map[*BezierSegment]*BezierCurve)
			for beziercurve := range stage.BezierCurves {
				for _, beziersegment_ := range beziercurve.BezierSegments {
					res[beziersegment_] = beziercurve
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of BezierSegment
	case BezierSegment:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Vector2
	case Vector2:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case BezierCurve:
		res = "BezierCurve"
	case BezierSegment:
		res = "BezierSegment"
	case Vector2:
		res = "Vector2"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *BezierCurve:
		res = "BezierCurve"
	case *BezierSegment:
		res = "BezierSegment"
	case *Vector2:
		res = "Vector2"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case BezierCurve:
		res = []string{"Name", "BezierSegments"}
	case BezierSegment:
		res = []string{"Name", "Start", "ControlPointStart", "ControlPointEnd", "End"}
	case Vector2:
		res = []string{"Name", "X", "Y"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case BezierCurve:
		var rf ReverseField
		_ = rf
	case BezierSegment:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BezierCurve"
		rf.Fieldname = "BezierSegments"
		res = append(res, rf)
	case Vector2:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *BezierCurve:
		res = []string{"Name", "BezierSegments"}
	case *BezierSegment:
		res = []string{"Name", "Start", "ControlPointStart", "ControlPointEnd", "End"}
	case *Vector2:
		res = []string{"Name", "X", "Y"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *BezierCurve:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "BezierSegments":
			for idx, __instance__ := range inferedInstance.BezierSegments {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *BezierSegment:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Start":
			if inferedInstance.Start != nil {
				res = inferedInstance.Start.Name
			}
		case "ControlPointStart":
			if inferedInstance.ControlPointStart != nil {
				res = inferedInstance.ControlPointStart.Name
			}
		case "ControlPointEnd":
			if inferedInstance.ControlPointEnd != nil {
				res = inferedInstance.ControlPointEnd.Name
			}
		case "End":
			if inferedInstance.End != nil {
				res = inferedInstance.End.Name
			}
		}
	case *Vector2:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case BezierCurve:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "BezierSegments":
			for idx, __instance__ := range inferedInstance.BezierSegments {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case BezierSegment:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Start":
			if inferedInstance.Start != nil {
				res = inferedInstance.Start.Name
			}
		case "ControlPointStart":
			if inferedInstance.ControlPointStart != nil {
				res = inferedInstance.ControlPointStart.Name
			}
		case "ControlPointEnd":
			if inferedInstance.ControlPointEnd != nil {
				res = inferedInstance.ControlPointEnd.Name
			}
		case "End":
			if inferedInstance.End != nil {
				res = inferedInstance.End.Name
			}
		}
	case Vector2:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
