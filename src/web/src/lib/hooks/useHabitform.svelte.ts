import type { Habit } from '$lib/habit/types';
import { validateHabit } from '$lib/habit/validation';

export function useHabitForm(habit: Habit | null = null) {
	let name = $state('');
	let description = $state('');
	let target = $state(0);
	let unit = $state('');
	let frequency = $state('');
	let aggregation = $state<'avg' | 'sum' | 'latest'>('avg');

	function resetForm(h: Habit | null) {
		name = h?.name ?? '';
		description = h?.description ?? '';
		target = h?.target ?? 1;
		unit = h?.unit ?? 'time';
		frequency = h?.frequency ?? 'daily';
		aggregation = h?.aggregation ?? 'avg';
	}

	const formData = $derived.by(() => ({
		id: habit?.id ?? '',
		name,
		description,
		target: target === 0 ? null : Number(target),
		unit,
		frequency,
		aggregation
	}));

	const isValid = $derived.by(() => {
		return validateHabit(formData).isValid;
	});

	resetForm(habit);

	return {
		get name() {
			return name;
		},
		set name(v: string) {
			name = v;
		},

		get description() {
			return description;
		},
		set description(v: string) {
			description = v;
		},

		get target() {
			return target;
		},
		set target(v: number) {
			target = v;
		},

		get unit() {
			return unit;
		},
		set unit(v: string) {
			unit = v;
		},

		get frequency() {
			return frequency;
		},
		set frequency(v: string) {
			frequency = v;
		},

		get aggregation() {
			return aggregation;
		},
		set aggregation(v: 'avg' | 'sum' | 'latest') {
			aggregation = v;
		},

		// Form state
		get formData() {
			return formData;
		},
		get isValid() {
			return isValid;
		},

		resetForm
	};
}
