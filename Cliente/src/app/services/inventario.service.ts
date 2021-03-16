import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class InventarioService {
  API_URI = 'http://localhost:3000';

  constructor(private http: HttpClient) {
  }

  getInventario(id: number): Observable<any> {
    return this.http.get(`${this.API_URI}/inventario/${id}`);
  }
}
