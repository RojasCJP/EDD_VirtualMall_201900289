import {Component, OnInit} from '@angular/core';
import {TiendasService} from '../../services/tiendas.service';

@Component({
  selector: 'app-reportes',
  templateUrl: './reportes.component.html',
  styleUrls: ['./reportes.component.css']
})
export class ReportesComponent implements OnInit {

  constructor(private tiendasService: TiendasService) {
  }

  ngOnInit(): void {
    if (localStorage.getItem('arreglo').toString() !== 'true') {
      this.tiendasService.getArreglo().subscribe(
        res => {
          console.log('arreglo generado');
        },
        err => console.log(err)
      );
      this.tiendasService.getUsuariosGraph().subscribe(
        res => {
          console.log('usuarios generados');
        },
        err => console.log(err)
      );
      localStorage.setItem('arreglo', 'true');
    }
  }

}
