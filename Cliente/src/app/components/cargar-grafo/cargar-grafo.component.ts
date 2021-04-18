import {Component, OnInit} from '@angular/core';
import {Grafos} from "../../models/grafos";
import {TiendasService} from "../../services/tiendas.service";

@Component({
  selector: 'app-cargar-grafo',
  templateUrl: './cargar-grafo.component.html',
  styleUrls: ['./cargar-grafo.component.css']
})
export class CargarGrafoComponent implements OnInit {

  cuerpo: Grafos;
  cuerpoString: string;

  constructor(private grafoServises: TiendasService) {
  }

  ngOnInit(): void {
  }

  cargarGrafo(): void {
    this.cuerpo = JSON.parse(this.cuerpoString);
    console.log(this.cuerpo);
    this.grafoServises.cargarGrafos(this.cuerpo).subscribe(
      res => console.log(res),
      error => console.log(error)
    );
  }

}
