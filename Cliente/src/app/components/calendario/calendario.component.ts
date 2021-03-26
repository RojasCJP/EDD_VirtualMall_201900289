import {Component, OnInit} from '@angular/core';
import {CalendarioService} from '../../services/calendario.service';
import {Pedidos} from '../../models/pedidos';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-calendario',
  templateUrl: './calendario.component.html',
  styleUrls: ['./calendario.component.css']
})
export class CalendarioComponent implements OnInit {
  years: number[];
  months: number[];
  pedidos: Pedidos[];
  productos: number[];

  constructor(private calendarioService: CalendarioService, private activatedRoute: ActivatedRoute) {
  }

  ngOnInit(): void {
    const params = this.activatedRoute.snapshot.params;
    const year = params.year;
    const month = params.month;
    const day = params.day;
    const departament = params.departament;
    if (year === undefined && month === undefined && day === undefined && departament === undefined) {
      this.calendarioService.getYears().subscribe(
        res => {
          this.years = res;
          console.log(this.years);
        },
        error => console.log(error)
      )
      ;
    } else if (year !== undefined && month !== undefined && day !== undefined && departament !== undefined) {
      this.calendarioService.getProducts(year, month, day, departament).subscribe(
        res => {
          this.productos = res;
          console.log(this.productos);
        },
        error => console.log(error)
      );
    } else if (year !== undefined && month !== undefined) {
      this.calendarioService.getCalendar(year, month).subscribe(
        res => {
          this.pedidos = res;
          console.log(this.pedidos);
        },
        error => console.log(error)
      );
    } else if (year !== undefined) {
      this.calendarioService.getMonths(year).subscribe(
        res => {
          this.months = res;
          console.log(this.months);
        },
        error => console.log(error)
      )
      ;
    }
  }

  redierctYear(year: number): void {
    this.calendarioService.imageMonths(year).subscribe(
      res => console.log(res),
      error => console.log(error)
    );
    window.location.href = '/calendario/' + year;
  }

  redirectYearMonth(month: number): void {
    const params = this.activatedRoute.snapshot.params;
    const year = params.year;
    this.calendarioService.imageCalendar(year, month).subscribe(
      res => console.log(res),
      error => console.log(error));
    window.location.href = '/calendario/' + year + '/' + month;
  }

  redirectYearMonthDay(departament: string, day: number): void {
    const params = this.activatedRoute.snapshot.params
    const year = params.year;
    const month = params.month;
    this.calendarioService.imageProducts(year, month, day, departament).subscribe(
      res => console.log(res),
      error => console.log(error)
    );
    window.location.href = '/calendario/' + year + '/' + month + '/' + day + '/' + departament;
  }
}


