// generated code - do not edit

import { CountryAPI } from './country-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Hello } from './hello'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Country {

	static GONGSTRUCT_NAME = "Country"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Hello?: Hello

	AlternateHellos: Array<Hello> = []
}

export function CopyCountryToCountryAPI(country: Country, countryAPI: CountryAPI) {

	countryAPI.CreatedAt = country.CreatedAt
	countryAPI.DeletedAt = country.DeletedAt
	countryAPI.ID = country.ID

	// insertion point for basic fields copy operations
	countryAPI.Name = country.Name

	// insertion point for pointer fields encoding
	countryAPI.CountryPointersEncoding.HelloID.Valid = true
	if (country.Hello != undefined) {
		countryAPI.CountryPointersEncoding.HelloID.Int64 = country.Hello.ID  
	} else {
		countryAPI.CountryPointersEncoding.HelloID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	countryAPI.CountryPointersEncoding.AlternateHellos = []
	for (let _hello of country.AlternateHellos) {
		countryAPI.CountryPointersEncoding.AlternateHellos.push(_hello.ID)
	}

}

// CopyCountryAPIToCountry update basic, pointers and slice of pointers fields of country
// from respectively the basic fields and encoded fields of pointers and slices of pointers of countryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCountryAPIToCountry(countryAPI: CountryAPI, country: Country, frontRepo: FrontRepo) {

	country.CreatedAt = countryAPI.CreatedAt
	country.DeletedAt = countryAPI.DeletedAt
	country.ID = countryAPI.ID

	// insertion point for basic fields copy operations
	country.Name = countryAPI.Name

	// insertion point for pointer fields encoding
	country.Hello = frontRepo.map_ID_Hello.get(countryAPI.CountryPointersEncoding.HelloID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(countryAPI.CountryPointersEncoding.AlternateHellos)) {
		console.error('Rects is not an array:', countryAPI.CountryPointersEncoding.AlternateHellos);
		return;
	}

	country.AlternateHellos = new Array<Hello>()
	for (let _id of countryAPI.CountryPointersEncoding.AlternateHellos) {
		let _hello = frontRepo.map_ID_Hello.get(_id)
		if (_hello != undefined) {
			country.AlternateHellos.push(_hello!)
		}
	}
}
