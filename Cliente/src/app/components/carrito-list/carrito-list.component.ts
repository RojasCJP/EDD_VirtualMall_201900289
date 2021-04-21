import {Component, OnInit} from '@angular/core';
import {ElementoCarrito} from '../../models/elementoCarrito';
import {InventarioService} from '../../services/inventario.service';

@Component({
  selector: 'app-carrito-list',
  templateUrl: './carrito-list.component.html',
  styleUrls: ['./carrito-list.component.css']
})
export class CarritoListComponent implements OnInit {
  carrito: ElementoCarrito[];
  total = 0;

  constructor(private inventarioService: InventarioService) {
  }

  ngOnInit(): void {
    this.inventarioService.verCarrito().subscribe(
      res => {
        this.carrito = res;
        for (const producto of this.carrito) {
          this.total += (producto.PrecioProducto * producto.Cantidad);
        }
      },
      error => console.log(error)
    );
  }

  comprar(): void {
    this.inventarioService.comprar().subscribe(
      res => console.log(res),
      error => console.log(error)
    );
  }

  deleteElement(tienda: number, producto: number): void {
    this.inventarioService.deleteProduct(tienda, producto).subscribe(
      res => console.log(res),
      error => console.log(error)
    );
  }

  grafo(): void {
    this.inventarioService.graficarGrafo().subscribe(
      res => {
        console.log(res);
        localStorage.setItem('arreglo', 'false');
        window.location.href = '/reportes';
      },
      error => {
        console.log(error);
      }
    );
  }

}
