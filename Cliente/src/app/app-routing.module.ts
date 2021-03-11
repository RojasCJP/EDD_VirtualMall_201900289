import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {TiendaListComponent} from './components/tienda-list/tienda-list.component';
import {CarritoListComponent} from './components/carrito-list/carrito-list.component';
import {CalendarioComponent} from './components/calendario/calendario.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: '/',
    pathMatch: 'full'
  },
  {
    path: 'tiendas',
    component: TiendaListComponent
  },
  {
    path: 'carrito',
    component: CarritoListComponent
  },
  {
    path: 'calendario',
    component: CalendarioComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
