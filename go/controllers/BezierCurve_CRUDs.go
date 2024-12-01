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
var __BezierCurve__dummysDeclaration__ models.BezierCurve
var __BezierCurve_time__dummyDeclaration time.Duration

var mutexBezierCurve sync.Mutex

// An BezierCurveID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBezierCurve updateBezierCurve deleteBezierCurve
type BezierCurveID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BezierCurveInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBezierCurve updateBezierCurve
type BezierCurveInput struct {
	// The BezierCurve to submit or modify
	// in: body
	BezierCurve *orm.BezierCurveAPI
}

// GetBezierCurves
//
// swagger:route GET /beziercurves beziercurves getBezierCurves
//
// # Get all beziercurves
//
// Responses:
// default: genericError
//
//	200: beziercurveDBResponse
func (controller *Controller) GetBezierCurves(c *gin.Context) {

	// source slice
	var beziercurveDBs []orm.BezierCurveDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBezierCurves", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierCurve.GetDB()

	_, err := db.Find(&beziercurveDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	beziercurveAPIs := make([]orm.BezierCurveAPI, 0)

	// for each beziercurve, update fields from the database nullable fields
	for idx := range beziercurveDBs {
		beziercurveDB := &beziercurveDBs[idx]
		_ = beziercurveDB
		var beziercurveAPI orm.BezierCurveAPI

		// insertion point for updating fields
		beziercurveAPI.ID = beziercurveDB.ID
		beziercurveDB.CopyBasicFieldsToBezierCurve_WOP(&beziercurveAPI.BezierCurve_WOP)
		beziercurveAPI.BezierCurvePointersEncoding = beziercurveDB.BezierCurvePointersEncoding
		beziercurveAPIs = append(beziercurveAPIs, beziercurveAPI)
	}

	c.JSON(http.StatusOK, beziercurveAPIs)
}

// PostBezierCurve
//
// swagger:route POST /beziercurves beziercurves postBezierCurve
//
// Creates a beziercurve
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBezierCurve(c *gin.Context) {

	mutexBezierCurve.Lock()
	defer mutexBezierCurve.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBezierCurves", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierCurve.GetDB()

	// Validate input
	var input orm.BezierCurveAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create beziercurve
	beziercurveDB := orm.BezierCurveDB{}
	beziercurveDB.BezierCurvePointersEncoding = input.BezierCurvePointersEncoding
	beziercurveDB.CopyBasicFieldsFromBezierCurve_WOP(&input.BezierCurve_WOP)

	_, err = db.Create(&beziercurveDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBezierCurve.CheckoutPhaseOneInstance(&beziercurveDB)
	beziercurve := backRepo.BackRepoBezierCurve.Map_BezierCurveDBID_BezierCurvePtr[beziercurveDB.ID]

	if beziercurve != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), beziercurve)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, beziercurveDB)
}

// GetBezierCurve
//
// swagger:route GET /beziercurves/{ID} beziercurves getBezierCurve
//
// Gets the details for a beziercurve.
//
// Responses:
// default: genericError
//
//	200: beziercurveDBResponse
func (controller *Controller) GetBezierCurve(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBezierCurve", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierCurve.GetDB()

	// Get beziercurveDB in DB
	var beziercurveDB orm.BezierCurveDB
	if _, err := db.First(&beziercurveDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var beziercurveAPI orm.BezierCurveAPI
	beziercurveAPI.ID = beziercurveDB.ID
	beziercurveAPI.BezierCurvePointersEncoding = beziercurveDB.BezierCurvePointersEncoding
	beziercurveDB.CopyBasicFieldsToBezierCurve_WOP(&beziercurveAPI.BezierCurve_WOP)

	c.JSON(http.StatusOK, beziercurveAPI)
}

// UpdateBezierCurve
//
// swagger:route PATCH /beziercurves/{ID} beziercurves updateBezierCurve
//
// # Update a beziercurve
//
// Responses:
// default: genericError
//
//	200: beziercurveDBResponse
func (controller *Controller) UpdateBezierCurve(c *gin.Context) {

	mutexBezierCurve.Lock()
	defer mutexBezierCurve.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBezierCurve", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierCurve.GetDB()

	// Validate input
	var input orm.BezierCurveAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var beziercurveDB orm.BezierCurveDB

	// fetch the beziercurve
	_, err := db.First(&beziercurveDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	beziercurveDB.CopyBasicFieldsFromBezierCurve_WOP(&input.BezierCurve_WOP)
	beziercurveDB.BezierCurvePointersEncoding = input.BezierCurvePointersEncoding

	db, _ = db.Model(&beziercurveDB)
	_, err = db.Updates(&beziercurveDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	beziercurveNew := new(models.BezierCurve)
	beziercurveDB.CopyBasicFieldsToBezierCurve(beziercurveNew)

	// redeem pointers
	beziercurveDB.DecodePointers(backRepo, beziercurveNew)

	// get stage instance from DB instance, and call callback function
	beziercurveOld := backRepo.BackRepoBezierCurve.Map_BezierCurveDBID_BezierCurvePtr[beziercurveDB.ID]
	if beziercurveOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), beziercurveOld, beziercurveNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the beziercurveDB
	c.JSON(http.StatusOK, beziercurveDB)
}

// DeleteBezierCurve
//
// swagger:route DELETE /beziercurves/{ID} beziercurves deleteBezierCurve
//
// # Delete a beziercurve
//
// default: genericError
//
//	200: beziercurveDBResponse
func (controller *Controller) DeleteBezierCurve(c *gin.Context) {

	mutexBezierCurve.Lock()
	defer mutexBezierCurve.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBezierCurve", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongthree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBezierCurve.GetDB()

	// Get model if exist
	var beziercurveDB orm.BezierCurveDB
	if _, err := db.First(&beziercurveDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&beziercurveDB)

	// get an instance (not staged) from DB instance, and call callback function
	beziercurveDeleted := new(models.BezierCurve)
	beziercurveDB.CopyBasicFieldsToBezierCurve(beziercurveDeleted)

	// get stage instance from DB instance, and call callback function
	beziercurveStaged := backRepo.BackRepoBezierCurve.Map_BezierCurveDBID_BezierCurvePtr[beziercurveDB.ID]
	if beziercurveStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), beziercurveStaged, beziercurveDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
