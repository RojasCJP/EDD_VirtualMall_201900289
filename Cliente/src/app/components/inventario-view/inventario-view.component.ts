import {Component, OnInit} from '@angular/core';
import {InventarioElemento} from '../../models/inventarioElemento';
import {InventarioService} from '../../services/inventario.service';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-inventario-view',
  templateUrl: './inventario-view.component.html',
  styleUrls: ['./inventario-view.component.css']
})
export class InventarioViewComponent implements OnInit {

  inventario: InventarioElemento[] =
    [
      {
        Nombre: 'elemento1',
        Codigo: 1,
        Descripcion: 'descripcion1',
        Precio: 1,
        Cantidad: 1,
        Imagen: 'https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg'
      },
      {
        Nombre: 'elemento2',
        Codigo: 2,
        Descripcion: 'descripcion2',
        Precio: 2,
        Cantidad: 2,
        Imagen: 'https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg'
      },
      {
        Nombre: 'elemento2',
        Codigo: 2,
        Descripcion: 'descripcion2',
        Precio: 2,
        Cantidad: 2,
        Imagen: 'https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg'
      },
      {
        Nombre: 'elemento2',
        Codigo: 2,
        Descripcion: 'descripcion2',
        Precio: 2,
        Cantidad: 2,
        Imagen: 'https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg'
      }
    ];

  constructor(private inventarioService: InventarioService, private activatedRoute: ActivatedRoute) {
  }

  ngOnInit(): void {
    const params = this.activatedRoute.snapshot.params;
    if (params.id) {
      this.getListInventarios(params.id);
    }
    console.log(params);
  }

  getListInventarios(id: number): void {
    this.inventarioService.getList(id).subscribe(
      res => {
        this.inventario = res;
      },
      error => console.log(error));
  }

  addCarrito(producto: number): void {
    const params = this.activatedRoute.snapshot.params;
    const tienda = params.id;
    this.inventarioService.agregarCarrito(tienda, producto).subscribe(
      res => console.log(res),
      error => console.log(error)
    );
  }
}
