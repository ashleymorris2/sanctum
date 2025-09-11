<script lang="ts">
	import { Hash, ChartNoAxesCombined, Check, Component } from '@lucide/svelte';
	import DropdownButton from '../base/DropdownButton.svelte';

	let {
		id = $bindable(),
		onToggleDropdown,
		open = false
	} = $props<{
		id?: string | null;
		onToggleDropdown?: (id: string) => void;
		open?: boolean;
	}>();

	type Option = {
		value: 'avg' | 'num';
		label: string;
		description: string;
		icon: typeof Component;
	};

	const options = [
		{
			value: 'avg',
			label: 'Average',
			description: 'Keep an average over a time period',
			icon: ChartNoAxesCombined
		},
		{
			value: 'num',
			label: 'Count',
			description: 'How many times/units in a time period?',
			icon: Hash
		}
	] satisfies Option[];

	let selected = $state(options[0]);

	function select(option: (typeof options)[number]) {
		selected = option;
	}
</script>

<DropdownButton icon={selected.icon} buttonLabel={selected.label} {onToggleDropdown} {open}>
	{#snippet dropdownContent()}
		{#each options as option, i (i)}
			<button
				type="button"
				role="radio"
				aria-checked={selected.value === option.value}
				class="cursor-pointer rounded-sm border border-transparent bg-base-100 p-2 text-left transition hover:bg-base-300"
				onclick={() => select(option)}
			>
				<span class="flex w-full items-center justify-between">
					<span class="flex items-center gap-2">
						<span class="mr-2 opacity-50">
							<option.icon class="h-5 w-5" />
						</span>
						<span>
							<span class="text-sm font-bold">
								{option.label}
							</span>
							<br />
							<span class="text-sm opacity-60">
								{option.description}
							</span>
						</span>
					</span>
					<!-- Right check mark -->
					{#if selected.value === option.value}
						<Check class="mr-2 h-4 w-4 -translate-y-[1px] text-primary" />
					{/if}
				</span>
			</button>
		{/each}
	{/snippet}
</DropdownButton>
