import { AppAction } from "../models/app-state.model";
import { ApiKubernetesServiceForm } from "../../../../api";

export enum KubernetesActionType {
  GetDefaults = '[Kubernetes] Get Defaults',
  GetDefaultsSuccess = '[Kubernetes] Get Defaults Success'
}

const initialState: ApiKubernetesServiceForm = {
    serviceName: "",
    ports: [],
    variables: [],
    volumes: []
}

export function kubernetesReducer(state: ApiKubernetesServiceForm = initialState, action: AppAction): ApiKubernetesServiceForm {
    switch (action.type) {
        case KubernetesActionType.GetDefaults:
            return initialState;
        case KubernetesActionType.GetDefaultsSuccess:
            return action.payload;
        default:
            return state;
    }
}