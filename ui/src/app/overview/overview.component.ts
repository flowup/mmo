import { Component } from '@angular/core';
import { AppStateModel } from '../store/models/app-state.model';
import { Observable } from 'rxjs/Observable';
import { ApiService } from '../../../api/index';
import { Store } from '@ngrx/store';
import { ServiceActionType } from '../store/reducers/service.reducer';

@Component({
  selector: 'mmo-overview',
  templateUrl: './overview.component.html',
  styleUrls: ['./overview.component.scss']
})
export class OverviewComponent {

  displayedColumns = ['name', 'description'];
  servicesList$: Observable<ApiService[]>;

  constructor(private store: Store<AppStateModel>) {
    this.servicesList$ = this.store.select(appState => appState.services);
    this.store.dispatch({type: ServiceActionType.GetServices});
  }
}
