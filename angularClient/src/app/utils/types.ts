export type tasksType = {
  id: number;
  name: string;
  description: string;
  status: 'done' | 'pending' | 'cancelled';
  photo: string | null;
  hex_color: `#${string}`;
};
