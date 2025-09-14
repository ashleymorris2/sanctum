<script lang="ts">
	import type { Component } from '@lucide/svelte';
	import type { Snippet } from 'svelte';

	let {
		label,
		icon: DisplayIcon,
		children,
		class: className = '',
		selectElement = $bindable(null)
	} = $props<{
		label: string;
		icon: typeof Component;
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
				<span class="flex items-center opacity-50">
					<DisplayIcon class="h-4 w-4" />
				</span>
				<span class="leading-none font-semibold">{label}</span>
			</span>
			{@render children()}
		</label>
	</div>
{/snippet}

{#if selectElement}
	{@render content({ onclick: openSelect, onkeydown: openSelect, role: 'button', tabindex: '0' })}
{:else}
	{@render content()}
{/if}
