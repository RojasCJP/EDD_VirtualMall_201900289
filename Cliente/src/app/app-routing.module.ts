import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {TiendaListComponent} from './components/tienda-list/tienda-list.component';
import {CarritoListComponent} from './components/carrito-list/carrito-list.component';
import {CalendarioComponent} from './components/calendario/calendario.component';
import {TiendaEspecificaComponent} from './components/tienda-especifica/tienda-especifica.component';
import {CargarTiendasComponent} from './components/cargar-tiendas/cargar-tiendas.component';
import {EliminarTiendaComponent} from './components/eliminar-tienda/eliminar-tienda.component';
import {InventarioViewComponent} from './components/inventario-view/inventario-view.component';
import {InventarioFormComponent} from './components/inventario-form/inventario-form.component';

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
  },
  {
    path: 'tiendaEspecifica',
    component: TiendaEspecificaComponent
  },
  {
    path: 'cargarTienda',
    component: CargarTiendasComponent
  },
  {
    path: 'eliminarTienda',
    component: EliminarTiendaComponent
  },
  {
    path: 'inventarioView',
    component: InventarioViewComponent
  },
  {
    path: 'inventarioAdd',
    component: InventarioFormComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
