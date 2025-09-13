<script lang="ts">
	import { CircleCheckBig, Target } from '@lucide/svelte';
	import DropdownButton from '../base/DropdownButton.svelte';
	import DropdownInputField from '../input/DropdownInputField.svelte';

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
	let unitSelectElement = $state<HTMLSelectElement | null>(null);
	let frequencySelectElement = $state<HTMLSelectElement | null>(null);
</script>

<DropdownButton
	bind:id
	icon={CircleCheckBig}
	buttonLabel={label}
	contentClass="p-0"
	styles="w-auto min-w-[15em]"
	{onToggleDropdown}
	{open}
>
	{#snippet dropdownContent()}
		<DropdownInputField label="Target" icon={Target}>
			<input
				dir="ltr"
				id="target-input"
				class="text-right"
				type="number"
				bind:value={target}
				min="1"
				placeholder="15"
			/>
		</DropdownInputField>

		<DropdownInputField label="Unit" icon={Target} bind:selectElement={unitSelectElement}>
			<select
				class="input-simple select border-transparent bg-transparent text-right outline-0 focus:outline-transparent"
				bind:value={unit}
				bind:this={unitSelectElement}
				id="unit-input"
			>
				<option selected>Times</option>
				<option>Reps</option>
				<option>Amber</option>
				<option>Velvet</option>
			</select>
		</DropdownInputField>

		<DropdownInputField label="Frequency" icon={Target} bind:selectElement={frequencySelectElement}>
			<select
				class="input-simple select border-transparent bg-transparent text-right outline-0 focus:outline-transparent"
				bind:value={frequency}
				bind:this={frequencySelectElement}
				id="frequency-input"
			>
				<option selected>Daily</option>
				<option>Weekly</option>
			</select>
		</DropdownInputField>
	{/snippet}
</DropdownButton>
