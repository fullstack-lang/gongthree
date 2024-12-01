// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongthree/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	beziercurveDBs map[uint]*BezierCurveDB

	nextIDBezierCurveDB uint

	beziersegmentDBs map[uint]*BezierSegmentDB

	nextIDBezierSegmentDB uint

	vector2DBs map[uint]*Vector2DB

	nextIDVector2DB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		beziercurveDBs: make(map[uint]*BezierCurveDB),

		beziersegmentDBs: make(map[uint]*BezierSegmentDB),

		vector2DBs: make(map[uint]*Vector2DB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *BezierCurveDB:
		db.nextIDBezierCurveDB++
		v.ID = db.nextIDBezierCurveDB
		db.beziercurveDBs[v.ID] = v
	case *BezierSegmentDB:
		db.nextIDBezierSegmentDB++
		v.ID = db.nextIDBezierSegmentDB
		db.beziersegmentDBs[v.ID] = v
	case *Vector2DB:
		db.nextIDVector2DB++
		v.ID = db.nextIDVector2DB
		db.vector2DBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BezierCurveDB:
		delete(db.beziercurveDBs, v.ID)
	case *BezierSegmentDB:
		delete(db.beziersegmentDBs, v.ID)
	case *Vector2DB:
		delete(db.vector2DBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BezierCurveDB:
		db.beziercurveDBs[v.ID] = v
		return db, nil
	case *BezierSegmentDB:
		db.beziersegmentDBs[v.ID] = v
		return db, nil
	case *Vector2DB:
		db.vector2DBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BezierCurveDB:
		if existing, ok := db.beziercurveDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db BezierCurve github.com/fullstack-lang/gongthree/go, record not found")
		}
	case *BezierSegmentDB:
		if existing, ok := db.beziersegmentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db BezierSegment github.com/fullstack-lang/gongthree/go, record not found")
		}
	case *Vector2DB:
		if existing, ok := db.vector2DBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Vector2 github.com/fullstack-lang/gongthree/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]BezierCurveDB:
		*ptr = make([]BezierCurveDB, 0, len(db.beziercurveDBs))
		for _, v := range db.beziercurveDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]BezierSegmentDB:
		*ptr = make([]BezierSegmentDB, 0, len(db.beziersegmentDBs))
		for _, v := range db.beziersegmentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Vector2DB:
		*ptr = make([]Vector2DB, 0, len(db.vector2DBs))
		for _, v := range db.vector2DBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gongthree/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *BezierCurveDB:
		tmp, ok := db.beziercurveDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First BezierCurve Unkown entry %d", i))
		}

		beziercurveDB, _ := instanceDB.(*BezierCurveDB)
		*beziercurveDB = *tmp
		
	case *BezierSegmentDB:
		tmp, ok := db.beziersegmentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First BezierSegment Unkown entry %d", i))
		}

		beziersegmentDB, _ := instanceDB.(*BezierSegmentDB)
		*beziersegmentDB = *tmp
		
	case *Vector2DB:
		tmp, ok := db.vector2DBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Vector2 Unkown entry %d", i))
		}

		vector2DB, _ := instanceDB.(*Vector2DB)
		*vector2DB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, Unkown type")
	}
	
	return db, nil
}

