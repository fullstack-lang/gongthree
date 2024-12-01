// generated code - do not edit

//insertion point for imports
import { CountryAPI } from './country-api'

import { HelloAPI } from './hello-api'


export class BackRepoData {
	// insertion point for declarations
	CountryAPIs = new Array<CountryAPI>()

	HelloAPIs = new Array<HelloAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.CountryAPIs = data?.CountryAPIs || [];

		this.HelloAPIs = data?.HelloAPIs || [];

	}

}