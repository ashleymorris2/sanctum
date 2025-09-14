<script lang="ts">
	import { CircleCheckBig, Target } from '@lucide/svelte';
	import DropdownButton from '../base/DropdownButton.svelte';
	import DropdownInputField from '../base/DropdownInputField.svelte';
	import NumberPicker from '../base/NumberPicker.svelte';

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

	let unitSelectElement = $state<HTMLSelectElement | null>(null);
	let frequencySelectElement = $state<HTMLSelectElement | null>(null);

	let buttonLabel = $derived(`${target} ${unit} per ${frequency === 'daily' ? 'day' : 'week'}`);

	const unitOptions = [
		{ value: 'times', label: 'Times', selected: true },
		{ value: 'reps', label: 'Reps' },
		{ value: 'pages', label: 'Pages' }
	];

	const frequencyOptions = [
		{ value: 'daily', label: 'Daily', selected: true },
		{ value: 'weekly', label: 'Weekly' }
	];
</script>

<DropdownButton
	bind:id
	icon={CircleCheckBig}
	contentClass="p-0"
	styles="w-auto min-w-[18em]"
	{buttonLabel}
	{onToggleDropdown}
	{open}
>
	<DropdownInputField label="Target" icon={Target} class="px-1 py-0.5">
		<!-- <input
			dir="ltr"
			id="target-input"
			class="text-right"
			type="number"
			bind:value={target}
			min="1"
			placeholder="15"
		/> -->
		<NumberPicker />
	</DropdownInputField>

	<DropdownInputField
		label="Unit"
		icon={Target}
		class="px-1 py-0.5"
		bind:selectElement={unitSelectElement}
	>
		<select
			id="unit-input"
			class="input-simple select-simple cursor-pointer"
			bind:value={unit}
			bind:this={unitSelectElement}
		>
			{#each unitOptions as option}
				<option value={option.value} selected={option.selected}>{option.label}</option>
			{/each}
		</select>
	</DropdownInputField>

	<DropdownInputField
		label="Frequency"
		icon={Target}
		class="px-1 py-0.5"
		bind:selectElement={frequencySelectElement}
	>
		<select
			class="input-simple select-simple cursor-pointer"
			bind:value={frequency}
			bind:this={frequencySelectElement}
			id="frequency-input"
		>
			{#each frequencyOptions as option}
				<option value={option.value} selected={option.selected}>{option.label}</option>
			{/each}
		</select>
	</DropdownInputField>
</DropdownButton>
