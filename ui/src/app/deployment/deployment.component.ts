import { Component, OnInit } from '@angular/core';
import { AppStateModel, AppAction } from '../store/models/app-state.model';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { DeploymentActionType } from '../store/effects/deployment.effect';
import { ApiGithubDeployRequest, ApiKubernetesClusters, ApiKubernetesDeployRequest } from '../../../api';
import { KubernetesActionType } from '../store/reducers/kubernetes.reducer';
import { Actions } from '@ngrx/effects';
import { Subscription } from 'rxjs/Subscription';
import { MatDialog } from '@angular/material';
import { ConfirmDeployDialog } from './deployment.dialog';

@Component({
  selector: 'mmo-deployment',
  templateUrl: './deployment.component.html',
  styleUrls: ['./deployment.component.scss']
})
export class DeploymentComponent implements OnInit {

  environments: string[] = ["staging", "production"];

  environment: string = "";
  message: string = "";
  ref: string = "master";

  clusters$: Observable<ApiKubernetesClusters>;
  cluster: string = "";
  source: string = "staging";
  namespace: string = "default";

  success: Subscription;

  constructor(private store: Store<AppStateModel>, 
    private actions$: Actions,
    public dialog: MatDialog) { }

  ngOnInit() {
    this.store.dispatch({type: KubernetesActionType.LoadClusters});
    this.clusters$ = this.store.select((state) => (state.kubeClusters))
    this.success = this.actions$.ofType(DeploymentActionType.DeployKubeSuccess).subscribe((action: AppAction) => {
      let request: ApiKubernetesDeployRequest = {
        environment: this.source,
        cluster: this.cluster,
        namespace: this.namespace
      }

      let dialogRef = this.dialog.open(ConfirmDeployDialog, {
        width: '750px',
        data: {
          request: request,
          configs: action.payload
        }
      });
  
      dialogRef.afterClosed().subscribe(() => {
        console.log('The dialog was closed');
      });
      console.log(action.payload);
    })
  }

  deployGithub() {

    let request: ApiGithubDeployRequest = {
      environment: this.environment,
      message: this.message,
      ref: this.ref
    };

    this.store.dispatch({
      type: DeploymentActionType.DeployGithub,
      payload: request
    });
  }

  deployKubernetes() {
    let request: ApiKubernetesDeployRequest = {
      environment: this.source,
      cluster: this.cluster,
      namespace: this.namespace
    }

    this.store.dispatch({
      type: DeploymentActionType.DeployKube, 
      payload: request
    });
  }

  ngOnDestroy() {
    this.success.unsubscribe();
  }
}
