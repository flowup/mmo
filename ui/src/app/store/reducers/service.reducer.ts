import { AppAction } from "../models/app-state.model";
import { ApiService } from "../../../../api";

export enum ServiceActionType {
  GetServices = '[Service] Get Services',
  GetServicesSuccess = '[Service] Get Services Success'
}

export function serviceReducer(state: ApiService[] = [], action: AppAction): ApiService[] {
    switch (action.type) {
        case ServiceActionType.GetServices:
            return [];
        case ServiceActionType.GetServicesSuccess:
            return action.payload;
        default:
            return state;
    }
}