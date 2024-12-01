import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { BezierCurveAPI } from './beziercurve-api'
import { BezierCurve, CopyBezierCurveAPIToBezierCurve } from './beziercurve'
import { BezierCurveService } from './beziercurve.service'

import { BezierSegmentAPI } from './beziersegment-api'
import { BezierSegment, CopyBezierSegmentAPIToBezierSegment } from './beziersegment'
import { BezierSegmentService } from './beziersegment.service'

import { Vector2API } from './vector2-api'
import { Vector2, CopyVector2APIToVector2 } from './vector2'
import { Vector2Service } from './vector2.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gongthree/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_BezierCurves = new Array<BezierCurve>() // array of front instances
	map_ID_BezierCurve = new Map<number, BezierCurve>() // map of front instances

	array_BezierSegments = new Array<BezierSegment>() // array of front instances
	map_ID_BezierSegment = new Map<number, BezierSegment>() // map of front instances

	array_Vector2s = new Array<Vector2>() // array of front instances
	map_ID_Vector2 = new Map<number, Vector2>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'BezierCurve':
				return this.array_BezierCurves as unknown as Array<Type>
			case 'BezierSegment':
				return this.array_BezierSegments as unknown as Array<Type>
			case 'Vector2':
				return this.array_Vector2s as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'BezierCurve':
				return this.map_ID_BezierCurve as unknown as Map<number, Type>
			case 'BezierSegment':
				return this.map_ID_BezierSegment as unknown as Map<number, Type>
			case 'Vector2':
				return this.map_ID_Vector2 as unknown as Map<number, Type>
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
		private beziercurveService: BezierCurveService,
		private beziersegmentService: BezierSegmentService,
		private vector2Service: Vector2Service,
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
		Observable<BezierCurveAPI[]>,
		Observable<BezierSegmentAPI[]>,
		Observable<Vector2API[]>,
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
			this.beziercurveService.getBezierCurves(this.GONG__StackPath, this.frontRepo),
			this.beziersegmentService.getBezierSegments(this.GONG__StackPath, this.frontRepo),
			this.vector2Service.getVector2s(this.GONG__StackPath, this.frontRepo),
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
			this.beziercurveService.getBezierCurves(this.GONG__StackPath, this.frontRepo),
			this.beziersegmentService.getBezierSegments(this.GONG__StackPath, this.frontRepo),
			this.vector2Service.getVector2s(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						beziercurves_,
						beziersegments_,
						vector2s_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var beziercurves: BezierCurveAPI[]
						beziercurves = beziercurves_ as BezierCurveAPI[]
						var beziersegments: BezierSegmentAPI[]
						beziersegments = beziersegments_ as BezierSegmentAPI[]
						var vector2s: Vector2API[]
						vector2s = vector2s_ as Vector2API[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_BezierCurves = []
						this.frontRepo.map_ID_BezierCurve.clear()

						beziercurves.forEach(
							beziercurveAPI => {
								let beziercurve = new BezierCurve
								this.frontRepo.array_BezierCurves.push(beziercurve)
								this.frontRepo.map_ID_BezierCurve.set(beziercurveAPI.ID, beziercurve)
							}
						)

						// init the arrays
						this.frontRepo.array_BezierSegments = []
						this.frontRepo.map_ID_BezierSegment.clear()

						beziersegments.forEach(
							beziersegmentAPI => {
								let beziersegment = new BezierSegment
								this.frontRepo.array_BezierSegments.push(beziersegment)
								this.frontRepo.map_ID_BezierSegment.set(beziersegmentAPI.ID, beziersegment)
							}
						)

						// init the arrays
						this.frontRepo.array_Vector2s = []
						this.frontRepo.map_ID_Vector2.clear()

						vector2s.forEach(
							vector2API => {
								let vector2 = new Vector2
								this.frontRepo.array_Vector2s.push(vector2)
								this.frontRepo.map_ID_Vector2.set(vector2API.ID, vector2)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						beziercurves.forEach(
							beziercurveAPI => {
								let beziercurve = this.frontRepo.map_ID_BezierCurve.get(beziercurveAPI.ID)
								CopyBezierCurveAPIToBezierCurve(beziercurveAPI, beziercurve!, this.frontRepo)
							}
						)

						// fill up front objects
						beziersegments.forEach(
							beziersegmentAPI => {
								let beziersegment = this.frontRepo.map_ID_BezierSegment.get(beziersegmentAPI.ID)
								CopyBezierSegmentAPIToBezierSegment(beziersegmentAPI, beziersegment!, this.frontRepo)
							}
						)

						// fill up front objects
						vector2s.forEach(
							vector2API => {
								let vector2 = this.frontRepo.map_ID_Vector2.get(vector2API.ID)
								CopyVector2APIToVector2(vector2API, vector2!, this.frontRepo)
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
				frontRepo.array_BezierCurves = []
				frontRepo.map_ID_BezierCurve.clear()

				backRepoData.BezierCurveAPIs.forEach(
					beziercurveAPI => {
						let beziercurve = new BezierCurve
						frontRepo.array_BezierCurves.push(beziercurve)
						frontRepo.map_ID_BezierCurve.set(beziercurveAPI.ID, beziercurve)
					}
				)

				// init the arrays
				frontRepo.array_BezierSegments = []
				frontRepo.map_ID_BezierSegment.clear()

				backRepoData.BezierSegmentAPIs.forEach(
					beziersegmentAPI => {
						let beziersegment = new BezierSegment
						frontRepo.array_BezierSegments.push(beziersegment)
						frontRepo.map_ID_BezierSegment.set(beziersegmentAPI.ID, beziersegment)
					}
				)

				// init the arrays
				frontRepo.array_Vector2s = []
				frontRepo.map_ID_Vector2.clear()

				backRepoData.Vector2APIs.forEach(
					vector2API => {
						let vector2 = new Vector2
						frontRepo.array_Vector2s.push(vector2)
						frontRepo.map_ID_Vector2.set(vector2API.ID, vector2)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.BezierCurveAPIs.forEach(
					beziercurveAPI => {
						let beziercurve = frontRepo.map_ID_BezierCurve.get(beziercurveAPI.ID)
						CopyBezierCurveAPIToBezierCurve(beziercurveAPI, beziercurve!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BezierSegmentAPIs.forEach(
					beziersegmentAPI => {
						let beziersegment = frontRepo.map_ID_BezierSegment.get(beziersegmentAPI.ID)
						CopyBezierSegmentAPIToBezierSegment(beziersegmentAPI, beziersegment!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.Vector2APIs.forEach(
					vector2API => {
						let vector2 = frontRepo.map_ID_Vector2.get(vector2API.ID)
						CopyVector2APIToVector2(vector2API, vector2!, frontRepo)
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
export function getBezierCurveUniqueID(id: number): number {
	return 31 * id
}
export function getBezierSegmentUniqueID(id: number): number {
	return 37 * id
}
export function getVector2UniqueID(id: number): number {
	return 41 * id
}
