import { Injectable } from "@angular/core";
import { Actions, Effect } from "@ngrx/effects";
import { Observable } from 'rxjs/Observable';
import { ApiClientService } from "../../../../api";
import { AppAction } from "../models/app-state.model";
import { map, switchMap } from "rxjs/operators";

export enum DeploymentActionType {
    DeployGithub = '[Deployment] Deploy Github',
    DeployGithubSuccess = '[Deployment] Deploy Github Success'
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
}