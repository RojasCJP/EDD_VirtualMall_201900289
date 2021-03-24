import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CalendarioService {
  API_URL = 'http://localhost:3000';

  constructor(private http: HttpClient) {
  }

  getYears(): Observable<any> {
    return this.http.get(`${this.API_URL}/calendario`);
  }

  getMonths(year: number): Observable<any> {
    return this.http.get(`${this.API_URL}/calendario/${year}`);
  }

  getCalendar(year: number, month: number): Observable<any> {
    return this.http.get(`${this.API_URL}/calendario/${year}/${month}`);
  }

  getProducts(year: number, month: number, day: number, departament: string): Observable<any> {
    return this.http.get(`${this.API_URL}/calendario/${year}/${month}/${day}/${departament}`);
  }

  imageYears(): Observable<any> {
    return this.http.get(`${this.API_URL}/calendarioImage`);
  }

  imageMonths(year: number): Observable<any> {
    return this.http.get(`${this.API_URL}/calendarioImage/${year}`);
  }

  imageCalendar(year: number, month: number): Observable<any> {
    return this.http.get(`${this.API_URL}/calendarioImage/${year}/${month}`);
  }

  imageProducts(year: number, month: number, day: number, departament: string): Observable<any> {
    return this.http.get(`${this.API_URL}/calendarioImage/${year}/${month}/${day}/${departament}`);
  }
}
