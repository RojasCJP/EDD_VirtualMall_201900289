import {Component, OnInit} from '@angular/core';
import {Usuarios} from "../../models/usuarios";
import {TiendasService} from "../../services/tiendas.service";

@Component({
  selector: 'app-registro-usarios',
  templateUrl: './registro-usarios.component.html',
  styleUrls: ['./registro-usarios.component.css']
})
export class RegistroUsariosComponent implements OnInit {

  cuerpo: Usuarios;
  cuerpoString: string;

  constructor(private usuariosServices: TiendasService) {
  }

  ngOnInit(): void {
  }

  cargarUsuarios(): void {
    this.cuerpo = JSON.parse(this.cuerpoString);
    console.log(this.cuerpo);
    this.usuariosServices.cargarUsuarios(this.cuerpo).subscribe(
      res => {
        // console.log(res);
        window.location.href = 'login';
      },
      err => console.log(err)
    );
  }
}
