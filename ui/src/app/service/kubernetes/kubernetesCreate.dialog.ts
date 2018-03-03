import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { Observable } from 'rxjs/Observable';
import { ApiKubernetesServiceForm } from '../../../../api';

@Component({
    selector: 'kubernetes-create-dialog',
    templateUrl: 'kubernetesCreate.dialog.html',
})
export class KubernetesCreateDialog implements OnInit {
    
    form: Observable<ApiKubernetesServiceForm>;
    serviceID: string;

    constructor(
        public dialogRef: MatDialogRef<KubernetesCreateDialog>,
            @Inject(MAT_DIALOG_DATA) public data: any) {

            this.serviceID = data.serviceID;

            console.log(data);
    }

    ngOnInit(): void {
        
    }

    onNoClick(): void {
        this.dialogRef.close();
    }
}