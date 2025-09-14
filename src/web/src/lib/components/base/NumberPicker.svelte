<script lang="ts">
	import { Component, Icon } from '@lucide/svelte';
	import Minus from '@lucide/svelte/icons/minus';
	import Plus from '@lucide/svelte/icons/plus';

	// Make value a bindable prop with a default, declared as a prop
	let {
		value = $bindable(0),
		minValue = 0,
		maxValue = 10,
		class: className = ''
	} = $props<{ value: number; minValue: number; maxValue: number; class?: string }>();

	function increment() {
		if (value < maxValue) {
			value++;
		}
	}

	function decrement() {
		if (value > minValue) {
			value--;
		}
	}

	function handleInput(event: Event) {
		const input = event.target as HTMLInputElement;
		const newValue = Number.parseInt(input.value, 10);
		if (!Number.isNaN(newValue) && newValue >= minValue) {
			value = newValue;
		} else {
			input.value = value.toString();
		}
	}

	const uid = $props.id();
</script>

{#snippet controlButton(
	uid: string,
	id: string,
	ariaLabel: string,
	disabled: boolean,
	ButtonIcon: typeof Component,
	onClick: (e: MouseEvent) => void,
	className = '',
	props = {}
)}
	<button
		{id}
		onclick={onClick}
		class="flex h-8 w-12 items-center justify-center border-r border-l {className} border-base-content/30 bg-base-100 text-base-content/60 transition-colors hover:bg-base-300 hover:text-base-content disabled:cursor-not-allowed disabled:opacity-50"
		aria-label={ariaLabel}
		aria-labelledby="{id} {uid}"
		aria-controls={uid}
		{disabled}
		{...props}
	>
		<ButtonIcon size={14} aria-hidden="true" />
	</button>
{/snippet}

<div class="*:not-first:mt-2">
	<div
		class=" relative inline-flex h-7 {className ??
			'w-auto'} items-center overflow-hidden rounded-md border border-base-content/30 text-sm whitespace-nowrap focus-within:border focus-within:outline-hidden"
	>
		{@render controlButton(
			uid,
			'decrement-button',
			'Decrease value',
			value <= minValue,
			Minus,
			decrement,
			'border-l-transparent rounded-l-md'
		)}

		<input
			id={uid}
			type="text"
			bind:value
			oninput={handleInput}
			aria-labelledby={uid}
			autocomplete="off"
			inputmode="numeric"
			autocorrect="off"
			aria-roledescription="Number input"
			spellcheck="false"
			min={minValue}
			max={maxValue}
			class="text-foreground w-full grow bg-base-100 px-3 py-2 text-center text-sm font-medium tabular-nums focus:outline-hidden"
		/>

		{@render controlButton(
			uid,
			'increment-button',
			'Increase value',
			value >= maxValue,
			Plus,
			increment,
			'border-r-transparent rounded-r-md'
		)}
	</div>
</div>
