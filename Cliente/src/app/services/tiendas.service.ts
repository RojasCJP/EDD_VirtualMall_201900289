import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {TiendaEspecifica} from '../models/tiendaEspecifica';
import {Observable} from 'rxjs';

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

}
