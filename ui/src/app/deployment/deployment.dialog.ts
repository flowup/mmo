import { Component, Inject, OnInit, ViewChild } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { AceEditorComponent } from 'ng2-ace-editor';
import { ApiKubernetesConfigs, ApiKubernetesDeployRequest, ApiKubernetesConfig } from '../../../api';
import "brace/mode/sh";
import { AppStateModel, AppAction } from '../store/models/app-state.model';
import { Store } from '@ngrx/store';
import { DeploymentActionType } from '../store/effects/deployment.effect';
import { Actions } from '@ngrx/effects';
import { Subscription } from 'rxjs/Subscription';

@Component({
    selector: 'deployment-dialog',
    templateUrl: 'deployment.dialog.html',
})
export class ConfirmDeployDialog implements OnInit {

    @ViewChild('highlight') highlight: AceEditorComponent;
    displayedColumns = ["name", "path"];
    configs: ApiKubernetesConfig[] = [];

    namespace: string = "";
    cluster: string = "";

    request: ApiKubernetesDeployRequest;

    subscription: Subscription;

    logEnabled = false;
    log: string = "";

    constructor(
        public dialogRef: MatDialogRef<ConfirmDeployDialog>,
        private actions$: Actions,
        private store: Store<AppStateModel>,
            @Inject(MAT_DIALOG_DATA) public data: DeploymentDialogData) {
            this.configs = data.configs.configs;
            this.namespace = data.request.namespace;
            this.cluster = data.request.cluster;

            this.request = data.request;

            this.subscription = this.actions$.ofType(DeploymentActionType.ConfirmDeployKubeSuccess).subscribe(
                (action: AppAction) => {
                    this.logEnabled = true;
                    this.log = action.payload;
                    console.log(action.payload);
                }
            )
            
    }

    ngOnInit(): void {
    }

    onNoClick(): void {
        this.dialogRef.close();
    }

    ngOnDestroy(): void {
        this.subscription.unsubscribe();
    }

    deploy(): void {
        this.store.dispatch({type: DeploymentActionType.ConfirmDeployKube, payload: this.request});
    }
}

interface DeploymentDialogData {
    configs: ApiKubernetesConfigs
    request: ApiKubernetesDeployRequest
}