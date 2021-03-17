import {Component, OnInit} from '@angular/core';
import {InventarioEntrada} from '../../models/inventarioEntrada';
import {InventarioService} from '../../services/inventario.service';

@Component({
  selector: 'app-inventario-form',
  templateUrl: './inventario-form.component.html',
  styleUrls: ['./inventario-form.component.css']
})
export class InventarioFormComponent implements OnInit {
  cuerpo: InventarioEntrada;
  cuerpoString: string;

  constructor(private inventarioService: InventarioService) {
  }

  ngOnInit(): void {
  }

  cargarInventario(): void {
    this.cuerpo = JSON.parse(this.cuerpoString);
    console.log(this.cuerpo);
    this.inventarioService.addInventario(this.cuerpo).subscribe(
      res => {
        console.log(res);
        window.location.href = '/tiendas';
      },
      error => console.log(error)
    );
  }

}
