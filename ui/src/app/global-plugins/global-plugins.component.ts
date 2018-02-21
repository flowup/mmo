import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { ApiPlugin } from '../../../api';
import { Store } from '@ngrx/store';
import { AppStateModel } from '../store/models/app-state.model';
import { PluginActionType } from '../store/reducers/plugin.reducer';

@Component({
  selector: 'mmo-global-plugins',
  templateUrl: './global-plugins.component.html',
  styleUrls: ['./global-plugins.component.scss']
})
export class GlobalPluginsComponent implements OnInit {

  displayedColumns = ['name', 'version'];
  pluginList$: Observable<ApiPlugin[]>;

  constructor(private store: Store<AppStateModel>) {
    this.pluginList$ = this.store.select(appState => appState.plugins);
    this.store.dispatch({type: PluginActionType.GetPlugin});
  }

  ngOnInit() {
  }

}
