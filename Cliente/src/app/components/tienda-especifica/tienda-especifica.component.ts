import {Component, OnInit} from '@angular/core';
import {Tienda} from '../../models/tienda';
import {TiendasService} from '../../services/tiendas.service';
import {TiendaEspecifica} from '../../models/tiendaEspecifica';


@Component({
  selector: 'app-tienda-especifica',
  templateUrl: './tienda-especifica.component.html',
  styleUrls: ['./tienda-especifica.component.css']
})
export class TiendaEspecificaComponent implements OnInit {

  tienda: Tienda = {
    Nombre: '',
    Descripcion: '',
    Contacto: '',
    Calificacion: 0,
    Logo: '',
    IdTienda: 0
  };

  tiendaEspecifica: TiendaEspecifica = {
    Nombre: '',
    Departamento: '',
    Calificacion: 0
  };

  constructor(private tiendaService: TiendasService) {
  }

  ngOnInit(): void {
    // this.tiendaService.getTiendaEspecifica().subscribe(
    //   res => {
    //     this.tienda = res;
    //   },
    //   err => console.log(err)
    // );
  }

  consultTienda(): void {
    this.tiendaService.getTiendaEspecifica(this.tiendaEspecifica).subscribe(
      res => {
        console.log(res);
        this.tienda = res;
      },
      err => console.log(err)
    );
  }

}
