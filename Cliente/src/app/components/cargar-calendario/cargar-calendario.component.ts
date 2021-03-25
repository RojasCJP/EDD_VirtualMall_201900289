import {Component, OnInit} from '@angular/core';
import {CalendarioService} from '../../services/calendario.service';
import {Calendario} from '../../models/calendario';

@Component({
  selector: 'app-cargar-calendario',
  templateUrl: './cargar-calendario.component.html',
  styleUrls: ['./cargar-calendario.component.css']
})
export class CargarCalendarioComponent implements OnInit {
  cuerpo: Calendario;
  cuerpoString: string;

  constructor(private calendarioService: CalendarioService) {
  }

  ngOnInit(): void {
  }

  cargarCalendario(): void {
    this.cuerpo = JSON.parse(this.cuerpoString);
    console.log(this.cuerpo);
    this.calendarioService.cargarCalendario(this.cuerpo).subscribe(
      res => console.log(res),
      error => console.log(error)
    );
  }
}
