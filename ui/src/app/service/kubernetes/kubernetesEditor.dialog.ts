import { Component, Inject, OnInit, ViewChild } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { AceEditorComponent } from 'ng2-ace-editor';
import "brace/mode/yaml";
import { AppStateModel } from '../../store/models/app-state.model';
import { Store } from '@ngrx/store';
import { ServiceDetailActionType } from '../../store/reducers/serviceDetail.reducer';

@Component({
    selector: 'kubernetes-edit-dialog',
    templateUrl: 'kubernetesEditor.dialog.html',
})
export class KubernetesEditorDialog implements OnInit {

    @ViewChild('highlight') highlight: AceEditorComponent;
    
    constructor(
        public dialogRef: MatDialogRef<KubernetesEditorDialog>,
        private store: Store<AppStateModel>,
            @Inject(MAT_DIALOG_DATA) public data: any) {

            console.log(data);
    }

    ngOnInit(): void {
    }

    onNoClick(): void {
        this.dialogRef.close();
    }

    save(): void {
        console.log("saving");
        this.store.dispatch({type: ServiceDetailActionType.SaveKubernetesConfig, payload: this.data.config });
        this.dialogRef.close();
    }
}