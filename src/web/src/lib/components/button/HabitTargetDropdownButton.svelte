<script lang="ts">
	import { CircleCheckBig, Target, Scale, RulerDimensionLine, Repeat } from '@lucide/svelte';
	import DropdownButton from '../base/DropdownButton.svelte';
	import DropdownInputField from '../base/DropdownInputField.svelte';
	import NumberPicker from '../base/NumberPicker.svelte';
	import RadioPills from '../base/RadioPills.svelte';

	let {
		id,
		onToggleDropdown,
		target = $bindable(0),
		unit = $bindable(),
		frequency = $bindable('daily'),
		open = false
	} = $props<{
		id?: string;
		onToggleDropdown?: (id: string) => void;
		target?: number;
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
		{ value: 'pages', label: 'Pages' },
		{ value: 'steps', label: 'Steps' }
	];

	const frequencyOptions = [
		{ id: 'dly', value: 'daily', label: 'Daily' },
		{ id: 'wkly', value: 'weekly', label: 'Weekly' },
		{ id: 'cstm', value: 'custom', label: 'Custom' }
	];
</script>

<DropdownButton
	bind:id
	icon={CircleCheckBig}
	styles="w-auto min-w-[18em]"
	{buttonLabel}
	{onToggleDropdown}
	{open}
>
	<DropdownInputField label="Target">
		<NumberPicker bind:value={target} class="w-28" />
	</DropdownInputField>

	<DropdownInputField label="Unit" bind:selectElement={unitSelectElement}>
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

	<RadioPills legend="Frequency" options={frequencyOptions} bind:value={frequency} size="md" />
</DropdownButton>
