import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { CountryAPI } from './country-api'
import { Country, CopyCountryAPIToCountry } from './country'
import { CountryService } from './country.service'

import { HelloAPI } from './hello-api'
import { Hello, CopyHelloAPIToHello } from './hello'
import { HelloService } from './hello.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gongthree/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Countrys = new Array<Country>() // array of front instances
	map_ID_Country = new Map<number, Country>() // map of front instances

	array_Hellos = new Array<Hello>() // array of front instances
	map_ID_Hello = new Map<number, Hello>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Country':
				return this.array_Countrys as unknown as Array<Type>
			case 'Hello':
				return this.array_Hellos as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Country':
				return this.map_ID_Country as unknown as Map<number, Type>
			case 'Hello':
				return this.map_ID_Hello as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized");
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	GONG__StackPath: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	GONG__StackPath: string = ""
	private socket: WebSocket | undefined

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template 
		private countryService: CountryService,
		private helloService: HelloService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		);
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		);
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<CountryAPI[]>,
		Observable<HelloAPI[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
			// insertion point sub template
			this.countryService.getCountrys(this.GONG__StackPath, this.frontRepo),
			this.helloService.getHellos(this.GONG__StackPath, this.frontRepo),
		];

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.countryService.getCountrys(this.GONG__StackPath, this.frontRepo),
			this.helloService.getHellos(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						countrys_,
						hellos_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var countrys: CountryAPI[]
						countrys = countrys_ as CountryAPI[]
						var hellos: HelloAPI[]
						hellos = hellos_ as HelloAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Countrys = []
						this.frontRepo.map_ID_Country.clear()

						countrys.forEach(
							countryAPI => {
								let country = new Country
								this.frontRepo.array_Countrys.push(country)
								this.frontRepo.map_ID_Country.set(countryAPI.ID, country)
							}
						)

						// init the arrays
						this.frontRepo.array_Hellos = []
						this.frontRepo.map_ID_Hello.clear()

						hellos.forEach(
							helloAPI => {
								let hello = new Hello
								this.frontRepo.array_Hellos.push(hello)
								this.frontRepo.map_ID_Hello.set(helloAPI.ID, hello)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						countrys.forEach(
							countryAPI => {
								let country = this.frontRepo.map_ID_Country.get(countryAPI.ID)
								CopyCountryAPIToCountry(countryAPI, country!, this.frontRepo)
							}
						)

						// fill up front objects
						hellos.forEach(
							helloAPI => {
								let hello = this.frontRepo.map_ID_Hello.get(helloAPI.ID)
								CopyHelloAPIToHello(helloAPI, hello!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(GONG__StackPath: string): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath


		let params = new HttpParams().set("GONG__StackPath", this.GONG__StackPath)
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gongthree/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {


				const backRepoData = new BackRepoData(JSON.parse(event.data))

				let frontRepo = new (FrontRepo)

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_Countrys = []
				frontRepo.map_ID_Country.clear()

				backRepoData.CountryAPIs.forEach(
					countryAPI => {
						let country = new Country
						frontRepo.array_Countrys.push(country)
						frontRepo.map_ID_Country.set(countryAPI.ID, country)
					}
				)

				// init the arrays
				frontRepo.array_Hellos = []
				frontRepo.map_ID_Hello.clear()

				backRepoData.HelloAPIs.forEach(
					helloAPI => {
						let hello = new Hello
						frontRepo.array_Hellos.push(hello)
						frontRepo.map_ID_Hello.set(helloAPI.ID, hello)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.CountryAPIs.forEach(
					countryAPI => {
						let country = frontRepo.map_ID_Country.get(countryAPI.ID)
						CopyCountryAPIToCountry(countryAPI, country!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.HelloAPIs.forEach(
					helloAPI => {
						let hello = frontRepo.map_ID_Hello.get(helloAPI.ID)
						CopyHelloAPIToHello(helloAPI, hello!, frontRepo)
					}
				)



				observer.next(frontRepo)
			}
			this.socket!.onerror = event => {
				observer.error(event)
			}
			this.socket!.onclose = event => {
				observer.complete()
			}

			return () => {
				this.socket!.close()
			}
		})
	}
}

// insertion point for get unique ID per struct 
export function getCountryUniqueID(id: number): number {
	return 31 * id
}
export function getHelloUniqueID(id: number): number {
	return 37 * id
}