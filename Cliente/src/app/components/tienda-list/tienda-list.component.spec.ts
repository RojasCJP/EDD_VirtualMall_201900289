import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TiendaListComponent } from './tienda-list.component';

describe('TiendaListComponent', () => {
  let component: TiendaListComponent;
  let fixture: ComponentFixture<TiendaListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TiendaListComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TiendaListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
