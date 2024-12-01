// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	CountryAPIs []*CountryAPI

	HelloAPIs []*HelloAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, countryDB := range backRepo.BackRepoCountry.Map_CountryDBID_CountryDB {

		var countryAPI CountryAPI
		countryAPI.ID = countryDB.ID
		countryAPI.CountryPointersEncoding = countryDB.CountryPointersEncoding
		countryDB.CopyBasicFieldsToCountry_WOP(&countryAPI.Country_WOP)

		backRepoData.CountryAPIs = append(backRepoData.CountryAPIs, &countryAPI)
	}

	for _, helloDB := range backRepo.BackRepoHello.Map_HelloDBID_HelloDB {

		var helloAPI HelloAPI
		helloAPI.ID = helloDB.ID
		helloAPI.HelloPointersEncoding = helloDB.HelloPointersEncoding
		helloDB.CopyBasicFieldsToHello_WOP(&helloAPI.Hello_WOP)

		backRepoData.HelloAPIs = append(backRepoData.HelloAPIs, &helloAPI)
	}

}
