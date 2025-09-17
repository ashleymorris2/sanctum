<script lang="ts">
	type Option = { id: string; label: string; value: string; disabled?: boolean };

	let {
		legend = 'Choose one',
		name = `rg-${Math.random().toString(36).slice(2, 8)}`,
		options = [] as Option[],
		value = $bindable<string>(''),
		class: className = '',
		size = 'md' as 'sm' | 'md' | 'lg'
	} = $props();

	const pad = size === 'sm' ? 'px-3 py-1.5 text-sm' : size === 'lg' ? 'px-4 py-2.5' : 'px-4 py-2';
</script>

<fieldset class={`flex flex-col  ${className}`}>
	<div class="p-3">
		<legend class="mb-2 text-sm leading-none font-semibold select-none">{legend}</legend>

		<div class="flex flex-row gap-2">
			{#each options as option (option.id)}
				<!-- {#key option.id}
			
				{/key} -->

				<div>
					<input
						id={option.id}
						type="radio"
						{name}
						class={`peer sr-only`}
						value={option.value}
						bind:group={value}
						disabled={option.disabled}
						checked={true}
					/>

					<label
						for={option.id}
						class={`inline-flex flex-1 
					items-center gap-2 rounded-lg border border-base-content/20 bg-base-100 text-sm ${pad}
					cursor-pointer tracking-wide peer-checked:border-base-content/60 peer-checked:bg-base-200 peer-checked:text-base-content peer-checked:shadow-sm
					`}
					>
						<span class="leading-none">{option.label}</span>
					</label>
				</div>
			{/each}
		</div>
	</div>
</fieldset>
