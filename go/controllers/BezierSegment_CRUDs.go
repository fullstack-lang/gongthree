// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongthree/go/models"
	"github.com/fullstack-lang/gongthree/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __BezierSegment__dummysDeclaration__ models.BezierSegment
var __BezierSegment_time__dummyDeclaration time.Duration

var mutexBezierSegment sync.Mutex

// An BezierSegmentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBezierSegment updateBezierSegment deleteBezierSegment
type BezierSegmentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BezierSegmentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBezierSegment updateBezierSegment
type BezierSegmentInput struct {
	// The BezierSegment to submit or modify
	// in: body
	BezierSegment *orm.BezierSegmentAPI
}

// GetBezierSegments
//
// swagger:route GET /beziersegments beziersegments getBezierSegments
//
// # Get all beziersegments
//
// Responses:
// default: genericError
//
//	200: beziersegmentDBResponse
func (controller *Controller) GetBezierSegments(c *gin.Context) {

	// source slice
	var beziersegmentDBs []orm.BezierSegmentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBezierSegments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierSegment.GetDB()

	_, err := db.Find(&beziersegmentDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	beziersegmentAPIs := make([]orm.BezierSegmentAPI, 0)

	// for each beziersegment, update fields from the database nullable fields
	for idx := range beziersegmentDBs {
		beziersegmentDB := &beziersegmentDBs[idx]
		_ = beziersegmentDB
		var beziersegmentAPI orm.BezierSegmentAPI

		// insertion point for updating fields
		beziersegmentAPI.ID = beziersegmentDB.ID
		beziersegmentDB.CopyBasicFieldsToBezierSegment_WOP(&beziersegmentAPI.BezierSegment_WOP)
		beziersegmentAPI.BezierSegmentPointersEncoding = beziersegmentDB.BezierSegmentPointersEncoding
		beziersegmentAPIs = append(beziersegmentAPIs, beziersegmentAPI)
	}

	c.JSON(http.StatusOK, beziersegmentAPIs)
}

// PostBezierSegment
//
// swagger:route POST /beziersegments beziersegments postBezierSegment
//
// Creates a beziersegment
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBezierSegment(c *gin.Context) {

	mutexBezierSegment.Lock()
	defer mutexBezierSegment.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBezierSegments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierSegment.GetDB()

	// Validate input
	var input orm.BezierSegmentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create beziersegment
	beziersegmentDB := orm.BezierSegmentDB{}
	beziersegmentDB.BezierSegmentPointersEncoding = input.BezierSegmentPointersEncoding
	beziersegmentDB.CopyBasicFieldsFromBezierSegment_WOP(&input.BezierSegment_WOP)

	_, err = db.Create(&beziersegmentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBezierSegment.CheckoutPhaseOneInstance(&beziersegmentDB)
	beziersegment := backRepo.BackRepoBezierSegment.Map_BezierSegmentDBID_BezierSegmentPtr[beziersegmentDB.ID]

	if beziersegment != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), beziersegment)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, beziersegmentDB)
}

// GetBezierSegment
//
// swagger:route GET /beziersegments/{ID} beziersegments getBezierSegment
//
// Gets the details for a beziersegment.
//
// Responses:
// default: genericError
//
//	200: beziersegmentDBResponse
func (controller *Controller) GetBezierSegment(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBezierSegment", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierSegment.GetDB()

	// Get beziersegmentDB in DB
	var beziersegmentDB orm.BezierSegmentDB
	if _, err := db.First(&beziersegmentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var beziersegmentAPI orm.BezierSegmentAPI
	beziersegmentAPI.ID = beziersegmentDB.ID
	beziersegmentAPI.BezierSegmentPointersEncoding = beziersegmentDB.BezierSegmentPointersEncoding
	beziersegmentDB.CopyBasicFieldsToBezierSegment_WOP(&beziersegmentAPI.BezierSegment_WOP)

	c.JSON(http.StatusOK, beziersegmentAPI)
}

// UpdateBezierSegment
//
// swagger:route PATCH /beziersegments/{ID} beziersegments updateBezierSegment
//
// # Update a beziersegment
//
// Responses:
// default: genericError
//
//	200: beziersegmentDBResponse
func (controller *Controller) UpdateBezierSegment(c *gin.Context) {

	mutexBezierSegment.Lock()
	defer mutexBezierSegment.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBezierSegment", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierSegment.GetDB()

	// Validate input
	var input orm.BezierSegmentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var beziersegmentDB orm.BezierSegmentDB

	// fetch the beziersegment
	_, err := db.First(&beziersegmentDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	beziersegmentDB.CopyBasicFieldsFromBezierSegment_WOP(&input.BezierSegment_WOP)
	beziersegmentDB.BezierSegmentPointersEncoding = input.BezierSegmentPointersEncoding

	db, _ = db.Model(&beziersegmentDB)
	_, err = db.Updates(&beziersegmentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	beziersegmentNew := new(models.BezierSegment)
	beziersegmentDB.CopyBasicFieldsToBezierSegment(beziersegmentNew)

	// redeem pointers
	beziersegmentDB.DecodePointers(backRepo, beziersegmentNew)

	// get stage instance from DB instance, and call callback function
	beziersegmentOld := backRepo.BackRepoBezierSegment.Map_BezierSegmentDBID_BezierSegmentPtr[beziersegmentDB.ID]
	if beziersegmentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), beziersegmentOld, beziersegmentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the beziersegmentDB
	c.JSON(http.StatusOK, beziersegmentDB)
}

// DeleteBezierSegment
//
// swagger:route DELETE /beziersegments/{ID} beziersegments deleteBezierSegment
//
// # Delete a beziersegment
//
// default: genericError
//
//	200: beziersegmentDBResponse
func (controller *Controller) DeleteBezierSegment(c *gin.Context) {

	mutexBezierSegment.Lock()
	defer mutexBezierSegment.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBezierSegment", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierSegment.GetDB()

	// Get model if exist
	var beziersegmentDB orm.BezierSegmentDB
	if _, err := db.First(&beziersegmentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&beziersegmentDB)

	// get an instance (not staged) from DB instance, and call callback function
	beziersegmentDeleted := new(models.BezierSegment)
	beziersegmentDB.CopyBasicFieldsToBezierSegment(beziersegmentDeleted)

	// get stage instance from DB instance, and call callback function
	beziersegmentStaged := backRepo.BackRepoBezierSegment.Map_BezierSegmentDBID_BezierSegmentPtr[beziersegmentDB.ID]
	if beziersegmentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), beziersegmentStaged, beziersegmentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
