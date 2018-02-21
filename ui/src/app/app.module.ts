import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { OverviewComponent } from './overview/overview.component';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';
import { EffectsModule } from '@ngrx/effects';
import { PluginEffect } from './store/effects/plugin.effect';
import { ServiceEffect } from './store/effects/service.effect';
import { StoreModule, ActionReducerMap } from '@ngrx/store';
import { pluginReducer } from './store/reducers/plugin.reducer';
import { serviceReducer } from './store/reducers/service.reducer';
import { ApiClientService } from '../../api/api-client-service';
import { HttpClientModule } from '@angular/common/http';
import { AppStateModel } from './store/models/app-state.model';
import { GlobalPluginsComponent } from './global-plugins/global-plugins.component';

const reducerMap: ActionReducerMap<AppStateModel> = {
  plugins: pluginReducer,
  services: serviceReducer,
};

@NgModule({
  declarations: [
    AppComponent,
    OverviewComponent,
    GlobalPluginsComponent
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatCardModule,
    MatIconModule,
    MatSidenavModule,
    MatTableModule,
    MatToolbarModule,
    EffectsModule.forRoot([
      PluginEffect,
      ServiceEffect
    ]),
    StoreModule.forRoot(reducerMap)
  ], 
  providers: [
    ApiClientService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
