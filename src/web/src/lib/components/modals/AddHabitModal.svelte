<script lang="ts">
	import Modal from '../base/Modal.svelte';
	import type { Habit } from '$lib/habit/types';
	import { DropdownGroup } from '$lib/stores/dropdownGroup.svelte';
	import TargetInputButton from '../button/TargetInputButton.svelte';
	import HabitTypeButton from '../button/HabitTypeButton.svelte';
	import { useHabitForm } from '$lib/hooks/useHabitform.svelte';

	let {
		open = $bindable(false),
		habit = null as Habit | null,
		onConfirm = function () {}
	} = $props();

	const group = new DropdownGroup();
	const form = useHabitForm(habit);
	const habitBtnId = `habit-btn`;
	const targetBtnId = `target-btn`;

	let habitOpen = $derived(!!habitBtnId && group.openId === habitBtnId);
	let targetOpen = $derived(!!targetBtnId && group.openId === targetBtnId);

	$effect(() => {
		if (open) form.resetForm(habit);
	});

	function confirm() {
		if (!form.isValid) return;
		onConfirm?.(form.formData);
		open = false;
	}

	function onClickOutside(node: HTMLElement, onOutside: () => void) {
		const handler = (e: Event) => {
			if (!node.contains(e.target as Node)) onOutside();
		};
		document.addEventListener('pointerdown', handler);
		return { destroy: () => document.removeEventListener('pointerdown', handler) };
	}
</script>

<Modal bind:open bind:isValid={form.isValid} onConfirm={confirm}>
	{#snippet content()}
		<div class="p-4">
			<div class="w-full">
				<input
					type="text"
					bind:value={form.name}
					placeholder="Name"
					class="w-full bg-transparent pt-2 pb-2 pl-2 text-xl font-semibold outline-none focus:ring-0"
				/>
				<input
					type="text"
					bind:value={form.description}
					placeholder="Description"
					class="txt-sm w-full bg-transparent pb-6 pl-2 outline-none focus:ring-0"
				/>
			</div>
			<div use:onClickOutside={group.close} class="flex items-center gap-4 pt-4">
				<div>
					<HabitTypeButton id={habitBtnId} open={habitOpen} onToggleDropdown={group.toggle} />
				</div>
				<div>
					<TargetInputButton
						id={targetBtnId}
						open={targetOpen}
						onToggleDropdown={group.toggle}
						bind:frequency={form.frequency}
						bind:target={form.target}
						bind:unit={form.unit}
					/>
				</div>
			</div>
		</div>
	{/snippet}
</Modal>
