import {Component, OnInit} from '@angular/core';
import {TiendasService} from '../../services/tiendas.service';

@Component({
  selector: 'app-merkle',
  templateUrl: './merkle.component.html',
  styleUrls: ['./merkle.component.css']
})
export class MerkleComponent implements OnInit {

  constructor(private merkleService: TiendasService) {
  }

  ngOnInit(): void {
    if (localStorage.getItem('merkle').toString() !== 'true') {
      this.merkleService.getMerkleUsuarios().subscribe(
        res => console.log('merkle usuarios generado'),
        err => console.log(err)
      );
      this.merkleService.getMerkleTiendas().subscribe(
        res => console.log('merkle tiendas generado'),
        err => console.log(err)
      );
      this.merkleService.getMerkleInventario().subscribe(
        res => console.log('merkle inventario generado'),
        err => console.log(err)
      );
      this.merkleService.getMerkleCalendario().subscribe(
        res => console.log('merkle calendario generado'),
        err => console.log(err)
      );
      localStorage.setItem('merkle', 'true');
    }
  }

}
