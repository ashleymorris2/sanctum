<script lang="ts">
	import { CircleCheckBig } from '@lucide/svelte';

	let { target = 0, unit = $bindable(), frequency = $bindable() } = $props();
</script>

<div class="dropdown">
	<div
		tabindex="0"
		role="button"
		class="btn border-base-content/20 align-middle font-normal btn-outline btn-sm hover:bg-primary-content/5"
	>
		<CircleCheckBig class="h-4 w-4" />
		<span class="text-sm">{target} {unit} per {frequency === 'daily' ? 'day' : 'week'}</span>
	</div>
	<div class="dropdown-content shadow-2xl">
		<div class="card bg-base-300 shadow-2xl card-xs">
			<div class="card-body p-4">
				<div class="flex gap-4">
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
				</div>
			</div>
		</div>
	</div>
</div>
