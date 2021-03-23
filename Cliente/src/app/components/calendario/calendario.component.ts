import {Component, OnInit} from '@angular/core';
import {TreeviewItem, TreeviewConfig} from 'ngx-treeview';


@Component({
  selector: 'app-calendario',
  templateUrl: './calendario.component.html',
  styleUrls: ['./calendario.component.css']
})
export class CalendarioComponent implements OnInit {
  items: TreeviewItem[];
  config = TreeviewConfig.create({
    hasFilter: true,
    hasCollapseExpand: true
  });

  constructor() {
  }

  ngOnInit(): void {
    const treeview = {
        text: 'IT',
        value: 1,
        children: [
          {
            text: 'Programming',
            value: 2,
            children: [
              {
                text: 'Frontend',
                value: 3,
                children: [
                  {text: 'Angular 1', value: 4, checked: false},
                  {text: 'Angular 2', value: 4, checked: false},
                  {text: 'ReactJS', value: 4, checked: false},
                ],
              },
              {
                text: 'Backend',
                value: 3,
                children: [
                  {text: 'C#', value: 4, checked: false},
                  {text: 'Java', value: 4, checked: false},
                  {text: 'Python', value: 4, checked: false},
                ],
              },
            ],
          },
          {
            text: 'Networking',
            value: 2,
            children: [
              {text: 'Internet', value: 3, checked: false},
              {text: 'Security', value: 3, checked: false},
            ],
          },
        ],
      }
    ;
    this.items = [new TreeviewItem(treeview)];
  }

  getCheckados(): void {
    const checkeados = this.items[0].getSelection();
    console.log(checkeados);
  }
}


