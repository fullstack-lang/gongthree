// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { NodeAPI } from './node-api'
import { Node, CopyNodeToNodeAPI } from './node'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { SVGIconAPI } from './svgicon-api'
import { ButtonAPI } from './button-api'

@Injectable({
  providedIn: 'root'
})
export class NodeService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  NodeServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private nodesUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.nodesUrl = origin + '/api/github.com/fullstack-lang/gongtree/go/v1/nodes';
  }

  /** GET nodes from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<NodeAPI[]> {
    return this.getNodes(GONG__StackPath, frontRepo)
  }
  getNodes(GONG__StackPath: string, frontRepo: FrontRepo): Observable<NodeAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<NodeAPI[]>(this.nodesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<NodeAPI[]>('getNodes', []))
      );
  }

  /** GET node by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<NodeAPI> {
    return this.getNode(id, GONG__StackPath, frontRepo)
  }
  getNode(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<NodeAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.nodesUrl}/${id}`;
    return this.http.get<NodeAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched node id=${id}`)),
      catchError(this.handleError<NodeAPI>(`getNode id=${id}`))
    );
  }

  // postFront copy node to a version with encoded pointers and post to the back
  postFront(node: Node, GONG__StackPath: string): Observable<NodeAPI> {
    let nodeAPI = new NodeAPI
    CopyNodeToNodeAPI(node, nodeAPI)
    const id = typeof nodeAPI === 'number' ? nodeAPI : nodeAPI.ID
    const url = `${this.nodesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<NodeAPI>(url, nodeAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<NodeAPI>('postNode'))
    );
  }
  
  /** POST: add a new node to the server */
  post(nodedb: NodeAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<NodeAPI> {
    return this.postNode(nodedb, GONG__StackPath, frontRepo)
  }
  postNode(nodedb: NodeAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<NodeAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<NodeAPI>(this.nodesUrl, nodedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted nodedb id=${nodedb.ID}`)
      }),
      catchError(this.handleError<NodeAPI>('postNode'))
    );
  }

  /** DELETE: delete the nodedb from the server */
  delete(nodedb: NodeAPI | number, GONG__StackPath: string): Observable<NodeAPI> {
    return this.deleteNode(nodedb, GONG__StackPath)
  }
  deleteNode(nodedb: NodeAPI | number, GONG__StackPath: string): Observable<NodeAPI> {
    const id = typeof nodedb === 'number' ? nodedb : nodedb.ID;
    const url = `${this.nodesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<NodeAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted nodedb id=${id}`)),
      catchError(this.handleError<NodeAPI>('deleteNode'))
    );
  }

  // updateFront copy node to a version with encoded pointers and update to the back
  updateFront(node: Node, GONG__StackPath: string): Observable<NodeAPI> {
    let nodeAPI = new NodeAPI
    CopyNodeToNodeAPI(node, nodeAPI)
    const id = typeof nodeAPI === 'number' ? nodeAPI : nodeAPI.ID
    const url = `${this.nodesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<NodeAPI>(url, nodeAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<NodeAPI>('updateNode'))
    );
  }

  /** PUT: update the nodedb on the server */
  update(nodedb: NodeAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<NodeAPI> {
    return this.updateNode(nodedb, GONG__StackPath, frontRepo)
  }
  updateNode(nodedb: NodeAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<NodeAPI> {
    const id = typeof nodedb === 'number' ? nodedb : nodedb.ID;
    const url = `${this.nodesUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<NodeAPI>(url, nodedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated nodedb id=${nodedb.ID}`)
      }),
      catchError(this.handleError<NodeAPI>('updateNode'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in NodeService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("NodeService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}