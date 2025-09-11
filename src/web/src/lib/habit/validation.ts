import type { Habit } from '$lib/habit/types';

export const FREQ = ['daily', 'weekly', 'monthly'] as const;

export type Frequency = (typeof FREQ)[number];

export type HabitErrors = Partial<{
	name: string;
	target: string;
	frequency: string;
}>;

export function validateHabit(input: Habit): { errors: HabitErrors; isValid: boolean } {
	const errors: HabitErrors = {};

	if (!input.name?.trim()) {
		errors.name = 'Name is required';
	}

	// if (input.target === null || input.target <= 0) {
	// 	errors.target = 'Target must be ≥ 0';
	// }

	// if (!FREQ.includes(input.frequency as Frequency)) {
	// 	errors.frequency = `Use one of: ${FREQ.join(', ')}`;
	// }

	return { errors, isValid: Object.keys(errors).length === 0 };
}
