import {Component, OnInit, OnDestroy} from '@angular/core';
import {InventarioElemento} from '../../models/inventarioElemento';
import {InventarioService} from '../../services/inventario.service';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-inventario-view',
  templateUrl: './inventario-view.component.html',
  styleUrls: ['./inventario-view.component.css']
})
export class InventarioViewComponent implements OnInit, OnDestroy {

  imagen;
  inventario: InventarioElemento[] = [];

  constructor(private inventarioService: InventarioService, private activatedRoute: ActivatedRoute) {
  }

  ngOnInit(): void {
    const params = this.activatedRoute.snapshot.params;
    this.imagen = sessionStorage.getItem('imagen');
    if (params.id) {
      if (this.imagen !== params.id) {
        this.inventarioService.getInventario(params.id).subscribe(
          res => console.log(res),
          error => console.log(error)
        );
        sessionStorage.setItem('imagen', params.id);
      }
      this.getListInventarios(params.id);
    }
  }

  getListInventarios(id: number): void {
    this.inventarioService.getList(id).subscribe(
      res => {
        this.inventario = res;
      },
      error => console.log(error));
    // this.inventarioService.getInventario(id).subscribe(
    //   res => {
    //     window.location.href = '/inventarioView/' + id;
    //   },
    //   error => console.log(error)
    // );
  }

  addCarrito(producto: number): void {
    const params = this.activatedRoute.snapshot.params;
    const tienda = params.id;
    this.inventarioService.agregarCarrito(tienda, producto).subscribe(
      res => console.log(res),
      error => console.log(error)
    );
  }

  ngOnDestroy(): void {
    const params = this.activatedRoute.snapshot.params;
    this.imagen = sessionStorage.setItem('imagen', params.id);
  }
}

