import { ApiPlugin } from '../../../../api/models/api-plugin.model';
import { Action } from '@ngrx/store';

export interface AppStateModel {
    plugins: ApiPlugin[];
}

export interface AppAction extends Action {
    payload?: any;
}