import { Actions, Effect } from '@ngrx/effects';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';
import { ApiClientService } from '../../../../api';
import { map, switchMap } from 'rxjs/operators';
import { AppAction } from '../models/app-state.model';
import { KubernetesActionType } from '../reducers/kubernetes.reducer';

@Injectable()
export class KubernetesEffect {
    constructor(private actions$: Actions, private apiClient: ApiClientService) { }

    @Effect() getKubernetes$: Observable<AppAction> = this.actions$
        .ofType(KubernetesActionType.GetDefaults)
        .pipe(
            switchMap((action: AppAction) => this.apiClient.kubernetesFormFromPlugins(action.payload, "").pipe(
                map((response) => ({type: KubernetesActionType.GetDefaultsSuccess, payload: response}))
            ))
        )

    @Effect() setKubernetes$: Observable<AppAction> = this.actions$
        .ofType(KubernetesActionType.CreateConfig)
        .pipe(
            switchMap((action: AppAction) => this.apiClient.kubernetesConfigFromForm(action.payload.id, action.payload.form).pipe(
                map(() => ({type: KubernetesActionType.CreateConfigSuccess, payload: action.payload.id}))
            ))
        )
}

@Injectable()
export class KubernetesClustersEffect {
    constructor(private actions$: Actions, private apiClient: ApiClientService) { }

    @Effect() getKubernetesClusters$: Observable<AppAction> = this.actions$
        .ofType(KubernetesActionType.LoadClusters)
        .pipe(
            switchMap(() => this.apiClient.getKubernetesClusters().pipe(
                map((response) => ({type: KubernetesActionType.LoadClustersSuccess, payload: response}))
            ))
        )
}