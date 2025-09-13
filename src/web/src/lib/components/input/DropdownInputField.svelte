<script lang="ts">
	import type { Component } from '@lucide/svelte';
	import type { Snippet } from 'svelte';

	let {
		label,
		icon: DisplayIcon,
		children,
		selectElement = $bindable(null)
	} = $props<{
		label: string;
		icon: typeof Component;
		children: Snippet;
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
	<div class="flex flex-1 items-center" {...props}>
		<label class="input-simple flex items-center justify-between gap-8 border-transparent">
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
