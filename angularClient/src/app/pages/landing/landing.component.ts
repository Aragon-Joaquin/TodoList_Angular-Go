import { Component, inject, OnInit } from '@angular/core';
import { QueryTasksService } from '../../services/query-tasks.service';
import { TaskStateService } from '../../services/task-state.service';
import { TasksComponent } from '../../components/tasks/tasks.component';
import { AngularComponent } from '../../components/svg/angular/angular.component';
import { GolangComponent } from '../../components/svg/golang/golang.component';
import { InputComponent } from '../../components/input/input.component';

@Component({
  selector: 'app-landing',
  imports: [TasksComponent, AngularComponent, GolangComponent, InputComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.css',
})
export class LandingComponent implements OnInit {
  fetchEverything = inject(QueryTasksService);
  tasks = inject(TaskStateService);

  ngOnInit(): void {
    const result = this.fetchEverything.makeQuery({ method: 'GET' });

    if (result instanceof Error) return this.tasks.showError(result);

    result.subscribe((res) => {
      console.log(res);
      this.tasks.addTasks(res);
    });
  }
}
