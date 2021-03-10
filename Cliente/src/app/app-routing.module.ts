import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {TiendaListComponent} from './components/tienda-list/tienda-list.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: '/',
    pathMatch: 'full'
  }, {
    path: 'tiendas',
    component: TiendaListComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
