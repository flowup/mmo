import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { OverviewComponent } from './overview/overview.component';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatDialogModule } from '@angular/material/dialog';
import { MatTableModule } from '@angular/material/table';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';
import { EffectsModule } from '@ngrx/effects';
import { PluginEffect } from './store/effects/plugin.effect';
import { ServiceEffect } from './store/effects/service.effect';
import { ServiceDetailEffect } from './store/effects/serviceDetail.effect';
import { StoreModule, ActionReducerMap } from '@ngrx/store';
import { kubernetesReducer } from './store/reducers/kubernetes.reducer';
import { pluginReducer } from './store/reducers/plugin.reducer';
import { serviceReducer } from './store/reducers/service.reducer';
import { serviceDetailReducer } from './store/reducers/serviceDetail.reducer';
import { ApiClientService } from '../../api/api-client-service';
import { HttpClientModule } from '@angular/common/http';
import { AppStateModel } from './store/models/app-state.model';
import { GlobalPluginsComponent } from './global-plugins/global-plugins.component';
import { ServiceComponent } from './service/service.component';
import { KubernetesCreateDialog } from './service/kubernetes/kubernetesCreate.dialog';
import { KubernetesFormComponent } from './service/kubernetes/kubernetes-form/kubernetes-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { KubernetesEffect } from './store/effects/kubernetes.effect';
import { KubernetesEditorDialog } from './service/kubernetes/kubernetesEditor.dialog';
import { AceEditorModule } from 'ng2-ace-editor';
import { MatCheckboxModule, MatInputModule, MatDividerModule, MatTooltipModule } from '@angular/material';
import { KubernetesDeployDialog } from './service/kubernetes/kubernetesDeploy.dialog';

const reducerMap: ActionReducerMap<AppStateModel> = {
  plugins: pluginReducer,
  services: serviceReducer,
  serviceDetails: serviceDetailReducer,
  kubernetesForm: kubernetesReducer,
};

@NgModule({
  declarations: [
    AppComponent,
    OverviewComponent,
    GlobalPluginsComponent,
    ServiceComponent,
    KubernetesCreateDialog,
    KubernetesDeployDialog,
    KubernetesEditorDialog,
    KubernetesFormComponent
  ],
  imports: [
    AceEditorModule,
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatCheckboxModule,
    MatDialogModule,
    MatDividerModule,
    MatCardModule,
    MatIconModule,
    MatInputModule,
    MatTooltipModule,
    MatSidenavModule,
    MatTableModule,
    MatToolbarModule,
    ReactiveFormsModule,
    EffectsModule.forRoot([
      KubernetesEffect,
      PluginEffect,
      ServiceEffect,
      ServiceDetailEffect
    ]),
    StoreModule.forRoot(reducerMap)
  ], 
  providers: [
    ApiClientService,
    {
      provide: 'domain',
      useValue: 'http://localhost:50080'
    },
  ],
  bootstrap: [AppComponent],
  entryComponents: [
    KubernetesCreateDialog,
    KubernetesEditorDialog,
    KubernetesDeployDialog
  ]
})
export class AppModule { }
