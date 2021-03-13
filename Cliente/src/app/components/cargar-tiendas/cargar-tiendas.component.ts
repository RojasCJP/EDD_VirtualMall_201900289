import {Component, OnInit} from '@angular/core';
import {Cuerpo} from '../../models/cuerpo';
import {TiendasService} from '../../services/tiendas.service';

@Component({
  selector: 'app-cargar-tiendas',
  templateUrl: './cargar-tiendas.component.html',
  styleUrls: ['./cargar-tiendas.component.css']
})
export class CargarTiendasComponent implements OnInit {

  cuerpo: Cuerpo;
  cuerpoString: string;

  constructor(private tiendaService: TiendasService) {
  }

  ngOnInit(): void {
  }

  cargarTiendas(): void {
    this.cuerpo = JSON.parse(this.cuerpoString);
    console.log(this.cuerpo);
    this.tiendaService.cargarTiendas(this.cuerpo).subscribe(
      res => {
        console.log(res);
        window.location.href = '/tiendas';
      },
      err => console.log(err)
    );
  }

}
