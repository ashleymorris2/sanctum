<script lang="ts">
	import { CircleCheckBig } from '@lucide/svelte';
	import DropdownButton from '../base/DropdownButton.svelte';

	let { 
		target = 0, 
		unit = $bindable(), 
		frequency = $bindable() 
	} = $props();
	
	let label = $derived(`${target} ${unit} per ${frequency === 'daily' ? 'day' : 'week'}`);
</script>

<DropdownButton icon={CircleCheckBig} buttonLabel={label} contentClass="flex gap-4">
	{#snippet dropdownContent()}
		<div class="flex w-24 flex-col">
			<label class="label" for="target-input">
				<span class="text-sm">Target</span>
			</label>
			<input
				id="target-input"
				type="number"
				bind:value={target}
				min="1"
				placeholder="15"
				class="input-simple"
			/>
		</div>
		<div class="flex flex-1 flex-col">
			<label class="label" for="unit-input">
				<span class="text-sm">Unit</span>
			</label>
			<input
				id="unit-input"
				type="text"
				bind:value={unit}
				placeholder="times"
				class="input-simple"
			/>
		</div>
		<div class="flex flex-1 flex-col">
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
		</div>
	{/snippet}
</DropdownButton>
