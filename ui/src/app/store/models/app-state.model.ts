import { ApiPlugin } from '../../../../api/models/api-plugin.model';
import { Action } from '@ngrx/store';
import { ApiService, ApiKubernetesConfig, ApiKubernetesServiceForm } from '../../../../api';

export interface AppStateModel {
    plugins: ApiPlugin[];
    services: ApiService[];
    serviceDetails: StoreState<AppServiceDetail>;

    kubernetesForm: ApiKubernetesServiceForm
}

export interface AppServiceDetail {
    plugins: ApiPlugin[];
    kubernetes: ApiKubernetesConfig[];
    meta: ApiService;
}

export interface AppAction extends Action {
    payload?: any;
}

export interface StoreState<T> {
    entities: { [id: string]: T };
}