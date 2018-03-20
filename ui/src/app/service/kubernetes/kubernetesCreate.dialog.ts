import { Component, Inject, OnInit, ViewChild } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { KubernetesFormComponent } from './kubernetes-form/kubernetes-form.component';
import { KubernetesActionType } from '../../store/reducers/kubernetes.reducer';
import { Store } from '@ngrx/store';
import { AppStateModel } from '../../store/models/app-state.model';

@Component({
  selector: 'kubernetes-create-dialog',
  templateUrl: 'kubernetesCreate.dialog.html',
})
export class KubernetesCreateDialog implements OnInit {

  serviceID: string;

  @ViewChild('kubernetesForm')
  kubernetesForm: KubernetesFormComponent;

  constructor(
    private store: Store<AppStateModel>,
    public dialogRef: MatDialogRef<KubernetesCreateDialog>,
    @Inject(MAT_DIALOG_DATA) public data: any) {

    this.serviceID = data.serviceID;
  }

  ngOnInit(): void {

  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  generate(): void {
    this.store.dispatch({
      type: KubernetesActionType.CreateConfig, payload: {
        id: this.serviceID,
        form: this.kubernetesForm.getFilledForm()
      }
    });

    this.dialogRef.close();
  }
}
