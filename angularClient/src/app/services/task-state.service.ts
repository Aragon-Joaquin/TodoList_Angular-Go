import { Injectable, signal, WritableSignal } from '@angular/core';
import { tasksType } from '../utils/types';

interface CustomError {
  name: string;
  message: string;
}

@Injectable({
  providedIn: 'root',
})
export class TaskStateService {
  taskState: WritableSignal<tasksType[] | []> = signal([]);
  errorState: WritableSignal<CustomError | null> = signal(null);

  showError(err: Error) {
    return this.errorState.set({
      name: err.name ?? 'Unknown name',
      message: err.message ?? 'Unknown message',
    });
  }

  addTasks(tasks: tasksType[] | tasksType) {
    return this.taskState.update((prevState) => [
      ...prevState,
      ...(Array.isArray(tasks) ? tasks : [tasks]),
    ]);
  }
}
