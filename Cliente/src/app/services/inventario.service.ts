import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {InventarioEntrada} from '../models/inventarioEntrada';

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

  addInventario(inventario: InventarioEntrada): Observable<any> {
    return this.http.post(`${this.API_URI}/inventario`, inventario);
  }

  getList(id: number): Observable<any> {
    return this.http.get(`${this.API_URI}/listaInventario/${id}`);
  }

  verCarrito(): Observable<any> {
    return this.http.get(`${this.API_URI}/verCarrito`);
  }

  agregarCarrito(tienda: number, producto: number): Observable<any> {
    return this.http.get(`${this.API_URI}/addCarrito/${tienda}/${producto}`);
  }

  comprar(): Observable<any> {
    return this.http.get(`${this.API_URI}/comprar`);
  }

  deleteProduct(tienda: number, producto: number): Observable<any> {
    return this.http.get(`${this.API_URI}/deleteCarrito/${tienda}/${producto}`);
  }

  graficarGrafo(): Observable<any> {
    return this.http.get(`${this.API_URI}/grafo/graficar`);
  }
}
