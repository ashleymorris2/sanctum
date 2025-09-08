<script lang="ts">
	import { Plus } from '@lucide/svelte';
	import AddHabitModal from './modals/AddHabitModal.svelte';
	import type { Habit } from '$lib/habit/types';

	let open = $state(false);

	function onConfirm(habit: Habit) {
		alert(`submitted ${habit.name}`);
	}
</script>

<button class="btn text-white btn-primary lg:btn-sm" onclick={() => (open = true)}>
	<Plus size="16" />
	Add habit
</button>

<AddHabitModal bind:open {onConfirm} />
