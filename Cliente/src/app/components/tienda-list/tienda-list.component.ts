import {Component, OnInit} from '@angular/core';
import {TiendasService} from '../../services/tiendas.service';
import {Tienda} from '../../models/tienda';
import {InventarioService} from '../../services/inventario.service';

@Component({
  selector: 'app-tienda-list',
  templateUrl: './tienda-list.component.html',
  styleUrls: ['./tienda-list.component.css']
})
export class TiendaListComponent implements OnInit {

  tiendas: Tienda[] = [];

  constructor(private tiendaService: TiendasService, private inventarioService: InventarioService) {
  }

  ngOnInit(): void {
    this.tiendaService.getTiendas().subscribe(
      res => {
        this.tiendas = res;
      },
      err => console.log(err)
    );
  }

  verInventario(id: number): void {
    this.inventarioService.getInventario(id).subscribe(
      res => {
        window.location.href = '/inventarioView';
      },
      error => console.log(error)
    );
  }
}

// todo tengo ya hecho que me muestre todas las tiendas
