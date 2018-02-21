import { NgModule } from '@angular/core';
import { RouterModule, Route } from '@angular/router';
import { OverviewComponent } from './overview/overview.component';
import { GlobalPluginsComponent } from './global-plugins/global-plugins.component';

interface NavigationItem extends Route {
  label?: string;
}

export const navItems: NavigationItem[] = [
  {
    path: 'overview',
    component: OverviewComponent,
    label: 'Overview',
    children: [
      // {path: 'landing', component: LandingComponent},
      // {
      //   path: 'platform',
      //   component: PlatformComponent,
      //   canActivate: [ AuthGuardService ],
      //   children: [
      //     {path: 'services', component: ProjectsComponent},
      //     {path: 'service/:id', component: ServiceDetailComponent},
      //     {path: 'settings', component: SettingsComponent},
      //     {path: '', redirectTo: 'services', pathMatch: 'full'},
      //   ]
      // },
      // {path: '**', redirectTo: 'landing'},
    ]
  },
  {
    path: 'plugins',
    component: GlobalPluginsComponent,
    label: 'Plugins',
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
