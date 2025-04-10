import { inject, Injectable } from '@angular/core';
import { tasksType } from '../utils/types';
import { HttpClient, HttpContext } from '@angular/common/http';
import { Observable } from 'rxjs';

interface makeQueryProps {
  method: 'GET' | 'POST' | 'DELETE' | 'PATCH';
  body?: FormData;
  parameter?: `/${string}`;
}

@Injectable({
  providedIn: 'root',
})
export class QueryTasksService {
  private url: string = 'http://localhost:3030/tasks';
  private http = inject(HttpClient);

  private returnParameter = (parameter: makeQueryProps['parameter']) =>
    parameter != null ? `${this.url}${parameter}` : this.url;

  //! i could use the normal fetch and do a hashmap for the methods instead of a switch statement
  makeQuery({
    method,
    body,
    parameter,
  }: makeQueryProps): Observable<tasksType | tasksType[]> | Error {
    const hasParameter = this.returnParameter(parameter);
    try {
      switch (method) {
        case 'GET':
          return this.http.get<tasksType>(hasParameter, {
            reportProgress: true,
            responseType: 'json',
          });
        case 'POST':
          return this.http.post<tasksType>(hasParameter, body ?? [], {
            responseType: 'json',
          });
        case 'DELETE':
          return this.http.delete<tasksType>(hasParameter, {
            responseType: 'json',
          });
        case 'PATCH':
          return this.http.patch<tasksType>(hasParameter, body ?? [], {
            responseType: 'json',
          });
        default:
          throw new Error('Method not defined.');
      }
    } catch (err) {
      console.log(err);
      if (err instanceof Error) return err;
      return new Error('Unknown Error.');
    }
  }
}
