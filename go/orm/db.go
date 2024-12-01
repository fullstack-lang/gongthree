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

	countryDBs map[uint]*CountryDB

	nextIDCountryDB uint

	helloDBs map[uint]*HelloDB

	nextIDHelloDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		countryDBs: make(map[uint]*CountryDB),

		helloDBs: make(map[uint]*HelloDB),
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
	case *CountryDB:
		db.nextIDCountryDB++
		v.ID = db.nextIDCountryDB
		db.countryDBs[v.ID] = v
	case *HelloDB:
		db.nextIDHelloDB++
		v.ID = db.nextIDHelloDB
		db.helloDBs[v.ID] = v
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
	case *CountryDB:
		delete(db.countryDBs, v.ID)
	case *HelloDB:
		delete(db.helloDBs, v.ID)
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
	case *CountryDB:
		db.countryDBs[v.ID] = v
		return db, nil
	case *HelloDB:
		db.helloDBs[v.ID] = v
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
	case *CountryDB:
		if existing, ok := db.countryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Country github.com/fullstack-lang/gongthree/go, record not found")
		}
	case *HelloDB:
		if existing, ok := db.helloDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Hello github.com/fullstack-lang/gongthree/go, record not found")
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
	case *[]CountryDB:
		*ptr = make([]CountryDB, 0, len(db.countryDBs))
		for _, v := range db.countryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]HelloDB:
		*ptr = make([]HelloDB, 0, len(db.helloDBs))
		for _, v := range db.helloDBs {
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
	case *CountryDB:
		tmp, ok := db.countryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Country Unkown entry %d", i))
		}

		countryDB, _ := instanceDB.(*CountryDB)
		*countryDB = *tmp
		
	case *HelloDB:
		tmp, ok := db.helloDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Hello Unkown entry %d", i))
		}

		helloDB, _ := instanceDB.(*HelloDB)
		*helloDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongthree/go, Unkown type")
	}
	
	return db, nil
}

