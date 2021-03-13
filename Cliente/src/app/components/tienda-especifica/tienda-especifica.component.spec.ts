import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TiendaEspecificaComponent } from './tienda-especifica.component';

describe('TiendaEspecificaComponent', () => {
  let component: TiendaEspecificaComponent;
  let fixture: ComponentFixture<TiendaEspecificaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TiendaEspecificaComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TiendaEspecificaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
