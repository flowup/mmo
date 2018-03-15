import { Component, OnInit, Input } from '@angular/core';
import { FormGroup, FormControl, FormArray, Validators } from '@angular/forms';
import { Subscription } from 'rxjs/Subscription';
import { Store } from '@ngrx/store';
import { AppStateModel } from '../../../store/models/app-state.model';
import { KubernetesActionType } from '../../../store/reducers/kubernetes.reducer';

@Component({
  selector: 'mmo-kubernetes-form',
  templateUrl: './kubernetes-form.component.html',
  styleUrls: ['./kubernetes-form.component.scss']
})
export class KubernetesFormComponent implements OnInit {

  kubernetesForm: FormGroup;

  constructor(private store: Store<AppStateModel>) {
  }

  @Input() serviceID: string;
  formSubscription: Subscription;

  ngOnInit() {
    this.store.dispatch({type: KubernetesActionType.GetDefaults, payload: this.serviceID});
    this.formSubscription = this.store.select((store) => store.kubernetesForm).subscribe(form => {
      form.ports = form.ports || [];
      form.variables = form.variables || [];
      form.volumes = form.volumes || [];

      // TODO: set value instead
      this.kubernetesForm = new FormGroup({
        'serviceName': new FormControl(form.serviceName, Validators.required),
        'ports': new FormArray(form.ports.map(port => this.buildPort(port.name, port.port))),
        'variables': new FormArray(form.variables.map(variable => this.buildVariable(variable.name, variable.value))),
        'volumes': new FormArray([])
      });
    });
  }

  public getFilledForm() {
    return this.kubernetesForm.value
  }


  buildPort(name: string, value: string): FormGroup {
    return new FormGroup({
      'name': new FormControl(name, Validators.required),
      'port': new FormControl(value, Validators.required)
    });
  }

  buildVariable(name: string, value: string): FormGroup {
    return new FormGroup({
      'name': new FormControl(name, Validators.required),
      'value': new FormControl(value, Validators.required)
    });
  }

  ngOnDestroy(): void {
    this.formSubscription.unsubscribe();
  }

  buildGceVolume() {
    return new FormGroup({
      'name': new FormControl('', Validators.required),
      'gceDisk': new FormControl('', Validators.required),
      'mountPath': new FormControl('', Validators.required)
    })
  }

  buildPersistentVolume() {
    return new FormGroup({
      'name': new FormControl('', Validators.required),
      'pvcSizeGB': new FormControl('', Validators.required | Validators.min(1)),
      'mountPath': new FormControl('', Validators.required)
    })
  }
}
