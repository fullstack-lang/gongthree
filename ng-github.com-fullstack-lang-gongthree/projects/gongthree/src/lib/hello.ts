// generated code - do not edit

import { HelloAPI } from './hello-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Hello {

	static GONGSTRUCT_NAME = "Hello"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyHelloToHelloAPI(hello: Hello, helloAPI: HelloAPI) {

	helloAPI.CreatedAt = hello.CreatedAt
	helloAPI.DeletedAt = hello.DeletedAt
	helloAPI.ID = hello.ID

	// insertion point for basic fields copy operations
	helloAPI.Name = hello.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyHelloAPIToHello update basic, pointers and slice of pointers fields of hello
// from respectively the basic fields and encoded fields of pointers and slices of pointers of helloAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyHelloAPIToHello(helloAPI: HelloAPI, hello: Hello, frontRepo: FrontRepo) {

	hello.CreatedAt = helloAPI.CreatedAt
	hello.DeletedAt = helloAPI.DeletedAt
	hello.ID = helloAPI.ID

	// insertion point for basic fields copy operations
	hello.Name = helloAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
