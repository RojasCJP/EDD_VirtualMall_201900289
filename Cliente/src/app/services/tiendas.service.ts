import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {TiendaEspecifica} from '../models/tiendaEspecifica';
import {Observable} from 'rxjs';
import {Cuerpo} from '../models/cuerpo';

@Injectable({
  providedIn: 'root'
})
export class TiendasService {

  API_URI = 'http://localhost:3000';

  constructor(private http: HttpClient) {
  }

  getTiendas(): Observable<any> {
    return this.http.get(`${this.API_URI}/todasTiendas`);
  }

  getTiendaEspecifica(tiendaEspecifica: TiendaEspecifica): any {
    return this.http.post(`${this.API_URI}/TiendaEspecifica`, tiendaEspecifica);
  }

  getTiendaLinealizada(id: string): any {
    return this.http.get(`${this.API_URI}/id/${id}`);
  }

  cargarTiendas(cuerpo: Cuerpo): Observable<any> {
    return this.http.post(`${this.API_URI}/cargartienda`, cuerpo);
  }

  deleteTienda(tiendaEspecifica: TiendaEspecifica): Observable<any> {
    const options = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
      body: tiendaEspecifica,
    };
    return this.http.delete(`${this.API_URI}/Eliminar`, options);
  }
}
