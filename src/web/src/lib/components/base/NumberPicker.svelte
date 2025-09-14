<script lang="ts">
	import Minus from '@lucide/svelte/icons/minus';
	import Plus from '@lucide/svelte/icons/plus';

	let value = $state(15);
	const minValue = 0;

	function increment() {
		value++;
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

<div class="*:not-first:mt-2">
	<div
		class=" relative inline-flex h-7 w-24 items-center overflow-hidden rounded-md border text-sm whitespace-nowrap focus-within:border focus-within:outline-hidden"
	>
		<button
			id="decrement-button"
			onclick={decrement}
			class="flex h-8 w-10 items-center justify-center rounded-l-md border-r border-l border-base-content/20 border-l-transparent bg-transparent text-base-content/60 transition-colors hover:bg-base-300 hover:text-base-content disabled:cursor-not-allowed disabled:opacity-50"
			aria-label="Decrease value"
			aria-labelledby="decrement-button {uid}"
			aria-controls={uid}
			disabled={value <= minValue}
		>
			<Minus size={14} aria-hidden="true" />
		</button>
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
			class="bg-background text-foreground w-full grow px-3 py-2 text-center tabular-nums focus:outline-hidden"
		/>
		<button
			id="increment-button"
			onclick={increment}
			class="flex h-8 w-10 items-center justify-center rounded-r-md border-r border-l border-base-content/20 border-r-transparent bg-transparent text-base-content/60 transition-colors hover:bg-base-content/5 hover:text-base-content disabled:cursor-not-allowed disabled:opacity-50"
			aria-label="Increase value"
			aria-labelledby="increment-button {uid}"
			aria-controls={uid}
		>
			<Plus size={14} aria-hidden="true" />
		</button>
	</div>
</div>
