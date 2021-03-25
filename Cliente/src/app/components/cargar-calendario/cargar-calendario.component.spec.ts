import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargarCalendarioComponent } from './cargar-calendario.component';

describe('CargarCalendarioComponent', () => {
  let component: CargarCalendarioComponent;
  let fixture: ComponentFixture<CargarCalendarioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargarCalendarioComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargarCalendarioComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
