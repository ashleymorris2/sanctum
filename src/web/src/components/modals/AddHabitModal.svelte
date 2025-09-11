<script lang="ts">
	import Modal from '../base/Modal.svelte';
	import type { Habit } from '$lib/habit/types';
	import { validateHabit } from '$lib/habit/validation';
	import { shortId } from '$lib/crypto/random';
	import { DropdownGroup } from '$lib/stores/dropdownGroup.svelte';
	import TargetInputButton from '../button/TargetInputButton.svelte';
	import HabitTypeButton from '../button/HabitTypeButton.svelte';

	let {
		open = $bindable(false),
		habit = null as Habit | null,
		onConfirm = function () {}
	} = $props();

	const group = new DropdownGroup();

	let name = $state('');
	let description = $state('');
	let target = $state('');
	let unit = $state('');
	let frequency = $state('');

	let habitBtnId = `habit-btn-${shortId()}`;
	let targetBtnId = `target-btn-${shortId()}`;
	let habitOpen = $derived(!!habitBtnId && group.openId === habitBtnId);
	let targetOpen = $derived(!!targetBtnId && group.openId === targetBtnId);

	function resetFrom(h: Habit | null) {
		name = h?.name ?? '';
		description = h?.description ?? '';
		target = h?.target?.toString() ?? '1';
		unit = h?.unit ?? 'time';
		frequency = h?.frequency ?? 'daily';
	}

	$effect(() => {
		if (open) resetFrom(habit);
	});

	const input = $derived.by(() => ({
		id: habit?.id,
		name,
		target: target === '' ? null : Number(target),
		unit,
		frequency
	}));

	let habitIsValid = $derived.by(() => {
		return validateHabit(input).isValid;
	});

	function confirm() {
		if (!habitIsValid) return;
		onConfirm?.(input);
		open = false;
	}
</script>

<Modal bind:open bind:isValid={habitIsValid} onConfirm={confirm}>
	{#snippet content()}
		<div class="p-4">
			<div class="w-full">
				<input
					type="text"
					bind:value={name}
					placeholder="Name"
					class="w-full bg-transparent pt-2 pb-2 pl-2 text-xl font-semibold outline-none focus:ring-0"
				/>
				<input
					type="text"
					bind:value={description}
					placeholder="Description"
					class="txt-sm w-full bg-transparent pb-6 pl-2 outline-none focus:ring-0"
				/>
			</div>
			<div class="flex items-center gap-4 pt-4">
				<div>
					<HabitTypeButton id={habitBtnId} open={habitOpen} onToggleDropdown={group.toggle} />
				</div>
				<div>
					<TargetInputButton
						id={targetBtnId}
						open={targetOpen}
						onToggleDropdown={group.toggle}
						bind:frequency
						bind:target
						bind:unit
					/>
				</div>
			</div>
		</div>
	{/snippet}
</Modal>
