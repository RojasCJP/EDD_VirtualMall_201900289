import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegistroUsariosComponent } from './registro-usarios.component';

describe('RegistroUsariosComponent', () => {
  let component: RegistroUsariosComponent;
  let fixture: ComponentFixture<RegistroUsariosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegistroUsariosComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RegistroUsariosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
