import { Component, OnInit } from '@angular/core';
import { AppStateModel } from '../store/models/app-state.model';
import { Store } from '@ngrx/store';
import { DeploymentActionType } from '../store/effects/deployment.effect';
import { ApiGithubDeployRequest } from '../../../api';

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

  constructor(private store: Store<AppStateModel>) { }

  ngOnInit() {
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

}
