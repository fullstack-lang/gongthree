// insertion point for imports
import { HelloAPI } from './hello-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CountryAPI {

	static GONGSTRUCT_NAME = "Country"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	CountryPointersEncoding: CountryPointersEncoding = new CountryPointersEncoding
}

export class CountryPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	HelloID: NullInt64 = new NullInt64 // if pointer is null, Hello.ID = 0

	AlternateHellos: number[] = []
}
