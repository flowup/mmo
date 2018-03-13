import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Store } from '@ngrx/store';
import { AppStateModel, AppServiceDetail } from '../store/models/app-state.model';
import { ServiceDetailActionType } from '../store/reducers/serviceDetail.reducer';
import { Subscription } from 'rxjs/Subscription';
import { MatDialog } from '@angular/material';
import { KubernetesCreateDialog } from './kubernetes/kubernetesCreate.dialog';
import { ApiKubernetesConfig } from '../../../api';
import { KubernetesEditorDialog } from './kubernetes/kubernetesEditor.dialog';
import { SelectionModel } from '@angular/cdk/collections';
import { KubernetesDeployDialog } from './kubernetes/kubernetesDeploy.dialog';

@Component({
  selector: 'mmo-service',
  templateUrl: './service.component.html',
  styleUrls: ['./service.component.scss']
})
export class ServiceComponent implements OnInit {
  
  subscription: Subscription;
  service: AppServiceDetail;

  serviceID: string;

  displayedColumnsPlugins = ['name', 'version'];
  displayedColumnsKubernetes = ['select', 'name', 'type', 'actions'];

  selection: SelectionModel<ApiKubernetesConfig>;
  
  constructor(private route: ActivatedRoute, private store: Store<AppStateModel>,
        public dialog: MatDialog) {
    this.route;
    this.service = {
      meta: {
        name: "",
        description: "",
      },
      kubernetes: [],
      plugins: [],
    };

    const initialSelection: ApiKubernetesConfig[] = [];
    const allowMultiSelect = true;
    this.selection = new SelectionModel<ApiKubernetesConfig>(allowMultiSelect, initialSelection);
  }

  ngOnInit() {
    this.route.params.subscribe(({id}) => {
      this.serviceID = id;
      this.store.dispatch({type: ServiceDetailActionType.GetServiceDetail, payload: id});
      this.subscription = this.store.select((store) => store.serviceDetails.entities[id])
        .subscribe(serviceDetail => {
          if (serviceDetail != null) {
            this.service = serviceDetail;
          }          
        });
    });
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }

  createKubernetesResource() {
    let dialogRef = this.dialog.open(KubernetesCreateDialog, {
      width: '750px',
      data: { serviceID: this.serviceID }
    });

    dialogRef.afterClosed().subscribe(() => {
      console.log('The dialog was closed');
    });
  }

  editKubernetes(config: ApiKubernetesConfig) {

    let dialogRef = this.dialog.open(KubernetesEditorDialog, {
      width: '750px',
      data: { config: config }
    });

    dialogRef.afterClosed().subscribe(() => {
      console.log('The dialog was closed');
    });
  }

  deleteKubernetes(config: ApiKubernetesConfig) {
    this.store.dispatch({type: ServiceDetailActionType.RemoveKubernetesConfig, payload: config});
  }

  deploySelected() {

    let dialogRef = this.dialog.open(KubernetesDeployDialog, {
      width: '750px',
      data: this.selection.selected
    });

    dialogRef.afterClosed().subscribe(() => {
      console.log('The dialog was closed');
    });
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.service.kubernetes.length;
    return numSelected == numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
        this.selection.clear() :
        this.service.kubernetes.filter(row => row.type !== "Invalid Kubernetes config").forEach(row => this.selection.select(row));
  }
}
