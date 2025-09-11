<script lang="ts">
	import { CircleCheckBig, Target } from '@lucide/svelte';

	import DropdownButton from '../base/DropdownButton.svelte';

	let { target = 0, unit = $bindable(), frequency = $bindable() } = $props();

	let label = $derived(`${target} ${unit} per ${frequency === 'daily' ? 'day' : 'week'}`);
</script>

<DropdownButton icon={CircleCheckBig} buttonLabel={label} contentClass="flex gap-4">
	{#snippet dropdownContent()}
		<div class="flex w-[28rem] flex-1 flex-row items-center">
			<label class="input flex w-full items-center justify-between gap-8" for="target-input">
				<span class="inline-flex items-center gap-1">
					<span class="flex items-center gap-1 opacity-70">
						<Target class="h-4 w-4" />
					</span>
					<span class="leading-none">Target</span>
				</span>
				<input
					dir="ltr"
					id="target-input"
					class="grow text-right"
					type="number"
					bind:value={target}
					min="1"
					placeholder="15"
				/>
			</label>
		</div>
		<div class="flex flex-1 flex-col">
			<!-- <label class="label" for="unit-input">
				<span class="text-sm">Unit</span>
			</label> -->
			<!--			<select class="input-simple select" bind:value={unit}>-->
			<!--				<option selected>Times</option>-->
			<!--				<option>Reps</option>-->
			<!--				<option>Amber</option>-->
			<!--				<option>Velvet</option>-->
			<!--			</select>-->
			<!-- <input
				id="unit-input"
				type="text"
				bind:value={unit}
				placeholder="unit"
				class="input-simple"
			/> -->
		</div>
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
