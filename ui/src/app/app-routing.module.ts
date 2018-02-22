import { NgModule } from '@angular/core';
import { RouterModule, Route } from '@angular/router';
import { OverviewComponent } from './overview/overview.component';
import { GlobalPluginsComponent } from './global-plugins/global-plugins.component';
import { ServiceComponent } from './service/service.component';

interface NavigationItem extends Route {
  label?: string;
}

export const navItems: NavigationItem[] = [
  {
    path: 'overview',
    component: OverviewComponent,
    label: 'Overview',
    children: []
  },
  {
    path: 'plugins',
    component: GlobalPluginsComponent,
    label: 'Plugins',
  },{
    path: 'overview/:id', 
    component: ServiceComponent
  },
  {
    path: '',
    redirectTo: 'overview',
    pathMatch: 'full'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(navItems)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
