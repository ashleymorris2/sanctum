export type Habit = {
	id: string;
	name: string;
	target: number | null;
	unit: string;
	frequency: string;
	aggregation: 'avg' | 'sum' | 'latest';
};
