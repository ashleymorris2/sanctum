<script lang="ts">
	import { CircleCheckBig, Target } from '@lucide/svelte';
	import DropdownButton from '../base/DropdownButton.svelte';

	let {
		id,
		onToggleDropdown,
		target = $bindable('0'),
		unit = $bindable(),
		frequency = $bindable(),
		open = false
	} = $props<{
		id?: string;
		onToggleDropdown?: (id: string) => void;
		target?: string;
		unit?: string;
		frequency?: string;
		open?: boolean;
	}>();

	let label = $derived(`${target} ${unit} per ${frequency === 'daily' ? 'day' : 'week'}`);
</script>

<DropdownButton
	bind:id
	icon={CircleCheckBig}
	buttonLabel={label}
	contentClass="flex gap-4"
	{onToggleDropdown}
	{open}
>
	{#snippet dropdownContent()}
		<div class="flex flex-1 items-center">
			<label
				class="input flex w-full items-center justify-between gap-8 input-ghost"
				for="target-input"
			>
				<span class="inline-flex items-center gap-2">
					<span class="flex items-center opacity-50">
						<Target class="h-4 w-4" />
					</span>
					<span class="leading-none font-semibold">Target</span>
				</span>
				<input
					dir="ltr"
					id="target-input"
					class="text-right"
					type="number"
					bind:value={target}
					min="1"
					placeholder="15"
				/>
			</label>
		</div>
		<!-- <div class="flex flex-1 flex-col">
			<label class="label" for="unit-input">
				<span class="text-sm">Unit</span>
			</label>
			<select class="input-simple select" bind:value={unit}>
				<option selected>Times</option>
				<option>Reps</option>
				<option>Amber</option>
				<option>Velvet</option>
			</select>
			<input
				id="unit-input"
				type="text"
				bind:value={unit}
				placeholder="unit"
				class="input-simple"
			/>
		</div> -->
		<!-- <div class="flex flex-1 flex-col">
			<label class="label" for="frequency-input">
				<span class="text-sm">Frequency</span>
			</label>
			<input
				id="frequency-input"
				type="text"
				bind:value={frequency}
				placeholder="daily"
				class="input-simple w-full"
			/>
		</div> -->
	{/snippet}
</DropdownButton>
