import { ApiPlugin } from "../../../../api/models/api-plugin.model";
import { AppAction } from "../models/app-state.model";

export enum PluginActionType {
  GetPlugin = '[Plugin] Get Plugin',
  GetPluginSuccess = '[Plugin] Get Plugin Success'
}

export function pluginReducer(state: ApiPlugin[] = [], action: AppAction): ApiPlugin[] {
    switch (action.type) {
        case PluginActionType.GetPlugin:
            return [];
        case PluginActionType.GetPluginSuccess:
            return action.payload;
        default:
            return state;
    }
}