import { Component, Inject, OnInit, ViewChild } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { AceEditorComponent } from 'ng2-ace-editor';
import { ApiKubernetesConfig } from '../../../../api';
import "brace/mode/sh";

@Component({
    selector: 'kubernetes-deploy-dialog',
    templateUrl: 'kubernetesDeploy.dialog.html',
})
export class KubernetesDeployDialog implements OnInit {

    @ViewChild('highlight') highlight: AceEditorComponent;
    
    deployCommand = "";
    constructor(
        public dialogRef: MatDialogRef<KubernetesDeployDialog>,
            @Inject(MAT_DIALOG_DATA) public data: ApiKubernetesConfig[]) {

            data.forEach(element => {
                this.deployCommand = "# Run these command in your console: \n\n"
                this.deployCommand += "kubectl apply -f " + element.path + "\n";
            });

            console.log(data);
    }

    ngOnInit(): void {
    }

    onNoClick(): void {
        this.dialogRef.close();
    }
}