import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargarGrafoComponent } from './cargar-grafo.component';

describe('CargarGrafoComponent', () => {
  let component: CargarGrafoComponent;
  let fixture: ComponentFixture<CargarGrafoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargarGrafoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargarGrafoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
