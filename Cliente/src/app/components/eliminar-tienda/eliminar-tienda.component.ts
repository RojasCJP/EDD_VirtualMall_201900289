import {Component, OnInit} from '@angular/core';
import {TiendaEspecifica} from '../../models/tiendaEspecifica';
import {Tienda} from '../../models/tienda';
import {TiendasService} from "../../services/tiendas.service";

@Component({
  selector: 'app-eliminar-tienda',
  templateUrl: './eliminar-tienda.component.html',
  styleUrls: ['./eliminar-tienda.component.css']
})
export class EliminarTiendaComponent implements OnInit {

  tiendaEspecifica: TiendaEspecifica = {
    Nombre: '',
    Departamento: '',
    Calificacion: 0
  };
  tienda: Tienda;

  constructor(private tiendaService: TiendasService) {
  }

  ngOnInit(): void {
  }

  deleteTienda(): any {
    this.tiendaService.deleteTienda(this.tiendaEspecifica).subscribe(
      res => {
        console.log(res);
      },
      err => console.log(err)
    );
  }

}
