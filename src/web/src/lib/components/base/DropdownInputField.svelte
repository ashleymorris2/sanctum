<script lang="ts">
	import type { Component } from '@lucide/svelte';
	import type { Snippet } from 'svelte';

	let {
		label,
		icon: DisplayIcon = null,
		children,
		class: className = '',
		selectElement = $bindable(null)
	} = $props<{
		label: string;
		icon?: typeof Component;
		children: Snippet;
		class?: string;
		selectElement?: HTMLSelectElement | null;
	}>();

	function openSelect() {
		if (selectElement) {
			selectElement.focus();
			// Modern approach
			if ('showPicker' in selectElement && typeof selectElement.showPicker === 'function') {
				try {
					selectElement.showPicker();
				} catch (e) {
					// Fallback to click if showPicker fails
					selectElement.click();
				}
			} else {
				selectElement.click();
			}
		}
	}
</script>

{#snippet content(props = {})}
	<div class="flex flex-1 items-center hover:rounded-lg hover:bg-base-300 {className}" {...props}>
		<label
			class="input-simple flex cursor-pointer items-center justify-between gap-8 border-transparent focus-within:outline-none"
		>
			<span class="inline-flex items-center gap-2">
				{#if DisplayIcon}
					<span class="flex items-center opacity-40">
						<DisplayIcon class="h-3 w-3" />
					</span>
				{/if}
				<span class="leading-none font-semibold text-base-content/60 select-none">{label}</span>
			</span>
			<span class="ml-auto flex items-center justify-end">
				{@render children()}
			</span>
		</label>
	</div>
{/snippet}

{#if selectElement}
	{@render content({ onclick: openSelect, onkeydown: openSelect, role: 'button', tabindex: '0' })}
{:else}
	{@render content()}
{/if}
