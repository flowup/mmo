import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { Store } from '@ngrx/store';
import { AppStateModel } from '../../store/models/app-state.model';
import { KubernetesActionType } from '../../store/reducers/kubernetes.reducer';

@Component({
    selector: 'kubernetes-create-dialog',
    templateUrl: 'kubernetesCreate.dialog.html',
})
export class KubernetesCreateDialog implements OnInit {

    
    constructor(
        public dialogRef: MatDialogRef<KubernetesCreateDialog>,
            private store: Store<AppStateModel>,
            @Inject(MAT_DIALOG_DATA) public data: any) {

            console.log(data);
    }

    ngOnInit(): void {
        this.store.dispatch({type: KubernetesActionType.GetDefaults, payload: this.data.serviceID})
        this.store.select((store) => store.kubernetesForm).subscribe(
            form => {
                console.log(form);
            }
        )
    }

    onNoClick(): void {
        this.dialogRef.close();
    }
}