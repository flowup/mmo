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
}