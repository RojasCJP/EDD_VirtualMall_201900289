import {Component, OnInit} from '@angular/core';
import {TiendasService} from '../../services/tiendas.service';
import {Tienda} from '../../models/tienda';

@Component({
  selector: 'app-tienda-list',
  templateUrl: './tienda-list.component.html',
  styleUrls: ['./tienda-list.component.css']
})
export class TiendaListComponent implements OnInit {

  tiendas: any = [];

  constructor(private tiendaService: TiendasService) {
  }

  ngOnInit(): void {
    this.tiendaService.getTiendas().subscribe(
      res => {
        this.tiendas = res;
      },
      err => console.log(err)
    );
  }

}
// todo tengo ya hecho que me muestre todas las tiendas
