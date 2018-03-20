import { Component } from '@angular/core';
import { navItems } from './app-routing.module';

@Component({
  selector: 'mmo-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'mmo';
  navItems = navItems.filter(item => item.label != null);

  constructor() {

  }
}
