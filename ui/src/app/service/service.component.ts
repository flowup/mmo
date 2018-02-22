import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Store } from '@ngrx/store';
import { AppStateModel, AppServiceDetail } from '../store/models/app-state.model';
import { ServiceDetailActionType } from '../store/reducers/serviceDetail.reducer';
import { Subscription } from 'rxjs/Subscription';

@Component({
  selector: 'mmo-service',
  templateUrl: './service.component.html',
  styleUrls: ['./service.component.scss']
})
export class ServiceComponent implements OnInit {
  
  subscription: Subscription;
  service: AppServiceDetail;

  displayedColumnsPlugins = ['name', 'version'];
  displayedColumnsKubernetes = ['name', 'type', 'actions'];

  constructor(private route: ActivatedRoute, private store: Store<AppStateModel>) {
    this.route;
    this.service = {
      meta: {
        name: "",
        description: "",
      },
      kubernetes: [],
      plugins: [],
    };
  }

  ngOnInit() {
    this.route.params.subscribe(({id}) => {
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
}
