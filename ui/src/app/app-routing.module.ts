import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { OverviewComponent } from './app/overview/overview.component';

const routes: Routes = [{
  path: '',
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
  ],
  component: OverviewComponent
}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
