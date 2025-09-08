<script lang="ts">
	import { Hash, ChartNoAxesCombined, Check, Component, LucideComponent } from '@lucide/svelte';
	import DropdownButton from '../base/DropdownButton.svelte';
	import type { SvelteComponent } from 'svelte';

	type Option = {
		value: 'average' | 'target';
		label: string;
		description: string;
		icon: typeof Component;
	};

	const iconSizeClass = 'w-5 h-5';
	const options = [
		{
			value: 'average',
			label: 'Average',
			description: 'Keep an average over a time period',
			icon: ChartNoAxesCombined // Reference to the Svelte icon component
		},
		{
			value: 'target',
			label: 'Count',
			description: 'How many times/units in a time period?',
			icon: Hash
		}
	] satisfies Option[];

	let selected = options[0];

	function select(v: (typeof options)[number]) {
		selected = v;
	}
</script>

<DropdownButton icon={selected.icon} buttonLabel={selected.label}>
	{#snippet dropdownContent()}
		{#each options as option, i (i)}
			<button
				type="button"
				role="radio"
				aria-checked={selected.value === option.value}
				class={`cursor-pointer rounded-sm border p-2 text-left transition ${
					selected.value === option.value
						? 'border-transparent bg-base-200 hover:bg-base-300'
						: 'border-transparent bg-base-100 hover:bg-base-300'
				}`}
				on:click={() => select(option)}
			>
				<span class="flex w-full items-center justify-between">
					<span class="flex items-center gap-2">
						<span class="mr-2 opacity-25">
							<svelte:component this={option.icon} class={iconSizeClass} />
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
					<!-- Right check -->
					{#if selected.value === option.value}
						<Check class="mr-2 h-4 w-4 text-primary" />
					{/if}
				</span>
			</button>
		{/each}
	{/snippet}
</DropdownButton>
