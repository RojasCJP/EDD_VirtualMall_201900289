import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {HttpClientModule} from '@angular/common/http';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {NavigationComponent} from './components/navigation/navigation.component';
import {TiendaListComponent} from './components/tienda-list/tienda-list.component';
import {CarritoListComponent} from './components/carrito-list/carrito-list.component';
import {CalendarioComponent} from './components/calendario/calendario.component';

@NgModule({
  declarations: [
    AppComponent,
    NavigationComponent,
    TiendaListComponent,
    CarritoListComponent,
    CalendarioComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
