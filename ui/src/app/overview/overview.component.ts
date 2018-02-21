import { Component } from '@angular/core';
import { AppStateModel } from '../store/models/app-state.model';
import { Observable } from 'rxjs/Observable';
import { ApiPlugin } from '../../../api/index';
import { Store } from '@ngrx/store';
import { PluginActionType } from '../store/reducers/plugin.reducer';

@Component({
  selector: 'mmo-overview',
  templateUrl: './overview.component.html',
  styleUrls: ['./overview.component.scss']
})
export class OverviewComponent {
  pluginList$: Observable<ApiPlugin[]>;

  constructor(private store: Store<AppStateModel>) {
    this.pluginList$ = this.store.select(appState => appState.plugins);
    this.store.dispatch({type: PluginActionType.GetPlugin});
  }
}
