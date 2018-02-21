
import { Actions, Effect } from '@ngrx/effects';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';
import { ApiClientService } from '../../../../api';
import { map, switchMap } from 'rxjs/operators';
import { AppAction } from '../models/app-state.model';
import { ServiceActionType } from '../reducers/service.reducer';

@Injectable()
export class ServiceEffect {
    constructor(private actions$: Actions, private apiClient: ApiClientService) { }

    @Effect() getServices$: Observable<AppAction> = this.actions$
        .ofType(ServiceActionType.GetServices)
        .pipe(
            switchMap(() => this.apiClient.getServices().pipe(
                map(({services}) => ({type: ServiceActionType.GetServicesSuccess, payload: services}))
            ))
        )
}