import { ApiPlugin } from '../../../../api/models/api-plugin.model';
import { Action } from '@ngrx/store';
import { ApiService } from '../../../../api';

export interface AppStateModel {
    plugins: ApiPlugin[];
    services: ApiService[];
}

export interface AppAction extends Action {
    payload?: any;
}