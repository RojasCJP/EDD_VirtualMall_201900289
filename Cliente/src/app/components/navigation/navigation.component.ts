import {Component, OnInit} from '@angular/core';
import {CalendarioService} from '../../services/calendario.service';
import {window} from "rxjs/operators";

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.css']
})
export class NavigationComponent implements OnInit {

  constructor(private calendarioService: CalendarioService) {
  }

  ngOnInit(): void {
  }

  hacerImagenYears(): void {
    this.calendarioService.imageYears().subscribe(
      res => console.log(res),
      error => console.log(error)
    );
  }

  redirigirReportes(): void {
    localStorage.setItem('arreglo', 'false');
    location.href = '/reportes';
  }

  redirigirMerkle(): void {
    localStorage.setItem('merkle', 'false');
    location.href = '/merkle';
  }
}
