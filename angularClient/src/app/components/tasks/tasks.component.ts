import { Component, Input } from '@angular/core';
import { tasksType } from '../../utils/types';

@Component({
  selector: 'app-tasks',
  imports: [],
  providers: [],
  templateUrl: './tasks.component.html',
  styleUrl: './tasks.component.css',
})
export class TasksComponent {
  @Input({ required: true }) task: tasksType | undefined;
}
