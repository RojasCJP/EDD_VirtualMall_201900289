import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule} from '@angular/forms';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {NavigationComponent} from './components/navigation/navigation.component';
import {TiendaListComponent} from './components/tienda-list/tienda-list.component';
import {CarritoListComponent} from './components/carrito-list/carrito-list.component';
import {CalendarioComponent} from './components/calendario/calendario.component';
import {TiendaEspecificaComponent} from './components/tienda-especifica/tienda-especifica.component';
import {CargarTiendasComponent} from './components/cargar-tiendas/cargar-tiendas.component';
import {EliminarTiendaComponent} from './components/eliminar-tienda/eliminar-tienda.component';
import { InventarioViewComponent } from './components/inventario-view/inventario-view.component';
import { InventarioFormComponent } from './components/inventario-form/inventario-form.component';
import {TreeviewModule} from 'ngx-treeview';
import { CargarCalendarioComponent } from './components/cargar-calendario/cargar-calendario.component';
import { LoginComponent } from './components/login/login.component';
import { RegistroUsariosComponent } from './components/registro-usarios/registro-usarios.component';
import { FormulariosComponent } from './components/formularios/formularios.component';
import { CargarGrafoComponent } from './components/cargar-grafo/cargar-grafo.component';
import { ReportesComponent } from './components/reportes/reportes.component';

@NgModule({
  declarations: [
    AppComponent,
    NavigationComponent,
    TiendaListComponent,
    CarritoListComponent,
    CalendarioComponent,
    TiendaEspecificaComponent,
    CargarTiendasComponent,
    EliminarTiendaComponent,
    InventarioViewComponent,
    InventarioFormComponent,
    CargarCalendarioComponent,
    LoginComponent,
    RegistroUsariosComponent,
    FormulariosComponent,
    CargarGrafoComponent,
    ReportesComponent
  ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        HttpClientModule,
        FormsModule,
        TreeviewModule.forRoot()
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}
