import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {TiendaEspecifica} from '../models/tiendaEspecifica';
import {Observable} from 'rxjs';
import {Cuerpo} from '../models/cuerpo';
import {Login, Confirmacion} from '../models/login';
import {Usuarios} from '../models/usuarios';
import {Grafos} from '../models/grafos';

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

  login(cuenta: Login): Observable<Confirmacion> {
    // @ts-ignore
    return this.http.post(`${this.API_URI}/usuarios/consultar`, cuenta);
  }

  cargarUsuarios(cuerpo: Usuarios): Observable<any> {
    return this.http.post(`${this.API_URI}/usuarios/cargar`, cuerpo);
  }

  cargarGrafos(cuerpo: Grafos): Observable<any> {
    return this.http.post(`${this.API_URI}/grafo/cargar`, cuerpo);
  }

  getArreglo(): Observable<any> {
    return this.http.get(`${this.API_URI}/getArreglo`);
  }

  getUsuariosGraph(): Observable<any> {
    return this.http.get(`${this.API_URI}/usuarios/graficar`);
  }

  getUsuariosEncriptados(): Observable<any> {
    return this.http.get(`${this.API_URI}/usuarios/encriptado`);
  }

}
