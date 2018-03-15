import { Actions, Effect } from '@ngrx/effects';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';
import { PluginActionType } from '../reducers/plugin.reducer';
import { ApiClientService } from '../../../../api';
import { map, switchMap } from 'rxjs/operators';
import { AppAction } from '../models/app-state.model';

@Injectable()
export class PluginEffect {
    constructor(private actions$: Actions, private apiClient: ApiClientService) { }

    @Effect() getPlugins$: Observable<AppAction> = this.actions$
        .ofType(PluginActionType.GetPlugin)
        .pipe(
            switchMap(() => this.apiClient.getGlobalPlugins().pipe(
                map(({plugins}) => ({type: PluginActionType.GetPluginSuccess, payload: plugins}))
            ))
        )
}