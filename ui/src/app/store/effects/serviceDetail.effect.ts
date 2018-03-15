import { Actions, Effect } from '@ngrx/effects';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';
import { ApiClientService } from '../../../../api';
import { map, switchMap } from 'rxjs/operators';
import 'rxjs/add/observable/of';
import { combineLatest } from 'rxjs/observable/combineLatest';
import { AppAction } from '../models/app-state.model';
import { ServiceDetailActionType } from '../reducers/serviceDetail.reducer';
import { KubernetesActionType } from '../reducers/kubernetes.reducer';

@Injectable()
export class ServiceDetailEffect {
    constructor(private actions$: Actions, private apiClient: ApiClientService) { }

    @Effect() getServiceDetails$: Observable<AppAction> = this.actions$
        .ofType(ServiceDetailActionType.GetServiceDetail)
        .pipe(
            switchMap((action: AppAction) => {
                const plugins = this.apiClient.getPlugins(action.payload, "");
                const kubernetes = this.apiClient.getKubernetesConfigs(action.payload, "");
                return combineLatest(plugins, kubernetes).pipe(
                    map(([{plugins}, {configs}]) => (
                        {
                            type: ServiceDetailActionType.GetServicePluginsSuccess, 
                            payload: {
                                id: action.payload,
                                plugins: plugins,
                                kubernetes: configs
                            }
                        }
                    )),
                )
                
            })
        )
    @Effect() refreshServiceDetail$: Observable<AppAction> = this.actions$
        .ofType(KubernetesActionType.CreateConfigSuccess)
        .pipe(
            switchMap((action: AppAction) => {
                return Observable.of({ type: ServiceDetailActionType.GetServiceDetail, payload: action.payload})
            })
        )

    @Effect() saveKubernetesConfig$: Observable<AppAction> = this.actions$
        .ofType(ServiceDetailActionType.SaveKubernetesConfig)
        .pipe(
            switchMap((action: AppAction) => {
                    return this.apiClient.saveKuberentesConfig(action.payload).pipe(
                        map(() => ({ type: ServiceDetailActionType.SaveKubernetesConfigSuccess }))
                )
            })
        )

    @Effect() removeKubernetesConfig$: Observable<AppAction> = this.actions$
        .ofType(ServiceDetailActionType.RemoveKubernetesConfig)
        .pipe(
            switchMap((action: AppAction) => {
                    return this.apiClient.removeKubernetesConfig(action.payload).pipe(
                        map(() => ({ type: ServiceDetailActionType.RemoveKubernetesConfigSuccess }))
                )
            })
        )
}