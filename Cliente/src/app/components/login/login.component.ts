import {Component, OnInit} from '@angular/core';
import {TiendasService} from '../../services/tiendas.service';
import {Login} from '../../models/login';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  cuenta: string;
  login: boolean;
  Dpi: number;
  datos: Login = {
    Dpi: null,
    Password: ''
  };

  constructor(private loginServise: TiendasService) {
  }

  ngOnInit(): void {
  }

  verificarLogin(): void {
    // this.datos.Dpi = parseInt(this.Dpi);
    this.loginServise.login(this.datos).subscribe(
      res => {
        if (res.Consulta === true) {
          if (res.Cuenta === 'Usuario') {
            window.location.href = '/tiendas';
          } else if (res.Cuenta === 'Admin') {
            window.location.href = '/cargar';
          }
        } else {
          window.location.href = '/registro';
        }
      },
      err => console.log(err)
    );
  }

}
