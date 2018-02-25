import { Component, Inject, OnInit, ViewChild } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { AceEditorComponent } from 'ng2-ace-editor';
import "brace/mode/yaml";

@Component({
    selector: 'kubernetes-edit-dialog',
    templateUrl: 'kubernetesEditor.dialog.html',
})
export class KubernetesEditorDialog implements OnInit {

    @ViewChild('highlight') highlight: AceEditorComponent;
    
    constructor(
        public dialogRef: MatDialogRef<KubernetesEditorDialog>,
            @Inject(MAT_DIALOG_DATA) public data: any) {

            console.log(data);
    }

    ngOnInit(): void {
    }

    onNoClick(): void {
        this.dialogRef.close();
    }
}