import { AppAction } from "../models/app-state.model";
import { ApiKubernetesServiceForm, ApiKubernetesClusters } from "../../../../api";

export enum KubernetesActionType {
  GetDefaults = '[Kubernetes] Get Defaults',
  GetDefaultsSuccess = '[Kubernetes] Get Defaults Success',
  CreateConfig = '[Kubernetes] Create config',
  CreateConfigSuccess = '[Kubernetes] Create config Success',
  LoadClusters = '[Kubernetes] Load clusters',
  LoadClustersSuccess = '[Kubernetes] Load clusters success'
}

const initialState: ApiKubernetesServiceForm = {
    serviceName: "",
    ports: [],
    variables: [],
    volumes: [],
    projectName: ""
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

export function kubernetesClustersReducer(state: ApiKubernetesClusters = { clusters: [] }, action: AppAction): ApiKubernetesClusters {
    switch (action.type) {
        case KubernetesActionType.LoadClusters:
            return { clusters: []};
        case KubernetesActionType.LoadClustersSuccess:
            return action.payload;
        default:
            return state;
    }
}
