<script module lang="ts">
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import AddHabitModal from '$lib/components/modals/HabitFormModal.svelte';

	const { Story } = defineMeta({
		title: 'Modals/AddHabitModal',
		component: AddHabitModal,
		parameters: {
			layout: 'fullscreen'
		},
		args: {
			open: true,
			title: 'Add Habit',
			habit: null
		},
		argTypes: {
			open: { control: 'boolean' },
			title: { control: 'text' },
			habit: { control: 'object' }
		}
	});
</script>

<Story name="Empty" let:args>
	<AddHabitModal {...args} />
</Story>¬

<Story
	name="Valid"
	let:args
	args={{
		habit: {
			id: '123',
			name: 'Pushups',
			target: 50,
			unit: 'reps',
			frequency: 'daily',
			aggregation: 'avg'
		}
	}}
>
	<AddHabitModal {...args} />
</Story>
