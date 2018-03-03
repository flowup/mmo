import { AppAction } from "../models/app-state.model";
import { ApiKubernetesServiceForm } from "../../../../api";

export enum KubernetesActionType {
  GetDefaults = '[Kubernetes] Get Defaults',
  GetDefaultsSuccess = '[Kubernetes] Get Defaults Success',
  CreateConfig = '[Kubernetes] Create config',
  CreateConfigSuccess = '[Kubernetes] Create config'
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
        case KubernetesActionType.CreateConfig:
            return initialState;
        case KubernetesActionType.CreateConfigSuccess:
            return initialState;
        default:
            return state;
    }
}