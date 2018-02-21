import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import { OverviewComponent } from './app/overview/overview.component';
import { MatToolbarModule } from '@angular/material/toolbar';
import { EffectsModule } from '@ngrx/effects';
import { PluginEffect } from './store/effects/plugin.effect';
import { StoreModule, ActionReducerMap } from '@ngrx/store';
import { pluginReducer } from './store/reducers/plugin.reducer';
import { ApiClientService } from '../../api/api-client-service';
import { HttpClientModule } from '@angular/common/http';
import { AppStateModel } from './store/models/app-state.model';

const reducerMap: ActionReducerMap<AppStateModel> = {
  plugins: pluginReducer
};

@NgModule({
  declarations: [
    AppComponent,
    OverviewComponent
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    MatToolbarModule,
    EffectsModule.forRoot([
      PluginEffect
    ]),
    StoreModule.forRoot(reducerMap)
  ], 
  providers: [
    ApiClientService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
