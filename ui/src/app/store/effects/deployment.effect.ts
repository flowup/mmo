import { Injectable } from "@angular/core";
import { Actions, Effect } from "@ngrx/effects";
import { Observable } from 'rxjs/Observable';
import { ApiClientService } from "../../../../api";
import { AppAction } from "../models/app-state.model";
import { map, switchMap } from "rxjs/operators";

export enum DeploymentActionType {
    DeployGithub = '[Deployment] Deploy Github',
    DeployGithubSuccess = '[Deployment] Deploy Github Success',
    DeployKube = '[Deployment] Deploy Kubernetes',
    DeployKubeSuccess = '[Deployment] Deploy Kubernetes Success',
    ConfirmDeployKube = '[Deployment] Confirm Deploy Kubernetes',
    ConfirmDeployKubeSuccess = '[Deployment] Confirm Deploy Kubernetes Success'
  }

@Injectable()
export class DeploymentEffect {
    constructor(private actions$: Actions, private apiClient: ApiClientService) { }

    @Effect() deployGithub$: Observable<AppAction> = this.actions$
        .ofType(DeploymentActionType.DeployGithub)
        .pipe(
            switchMap((action: AppAction) => this.apiClient.githubDeploy(action.payload).pipe(
                map(() => ({type: DeploymentActionType.DeployGithubSuccess }))
            ))
        )

    @Effect() deployKubernetes$: Observable<AppAction> = this.actions$
        .ofType(DeploymentActionType.DeployKube)
        .pipe(
            switchMap((action: AppAction) => this.apiClient.kubernetesDeploy(action.payload).pipe(
                map((services) => ({type: DeploymentActionType.DeployKubeSuccess, payload: services}))
            )
        )
    )

    @Effect() confirmDeployKubernetes$: Observable<AppAction> = this.actions$
        .ofType(DeploymentActionType.ConfirmDeployKube)
        .pipe(
            switchMap((action: AppAction) => this.apiClient.confirmKubernetesDeploy(action.payload).pipe(
                map((log) => ({type: DeploymentActionType.ConfirmDeployKubeSuccess, payload: log.output}))
            )
        )
    )
}