import { AppAction, AppServiceDetail, StoreState } from "../models/app-state.model";

export enum ServiceDetailActionType {
  GetServiceDetail = '[ServiceDetail] Get Service detail',
  GetServicePluginsSuccess = '[ServiceDetail] Get Service plugins Success',
  SaveKubernetesConfig = '[KubeConfig] Save Kubernetes config',
  SaveKubernetesConfigSuccess = '[KubeConfig] Save Kubernetes config success',
  RemoveKubernetesConfig = '[KubeConfig] Remove Kubernetes config',
  RemoveKubernetesConfigSuccess = '[KubeConfig] Remove Kubernetes config success',
}

const initialState: StoreState<AppServiceDetail> = {
    entities: {}
}

export function serviceDetailReducer(state: StoreState<AppServiceDetail> = initialState, action: AppAction): StoreState<AppServiceDetail> {
    switch (action.type) {
        case ServiceDetailActionType.GetServiceDetail:
            return state;
        case ServiceDetailActionType.GetServicePluginsSuccess:
            const entities = {
                ...state.entities,
                [action.payload.id]: {
                    plugins: action.payload.plugins,
                    kubernetes: action.payload.kubernetes,
                    meta: {
                        name: "",
                        description: ""
                    }
                }
            }
            
            return {
                ...state,
                entities
            };
        default:
            return state;
    }
}