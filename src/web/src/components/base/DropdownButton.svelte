<script lang="ts">
	import type { Snippet } from 'svelte';
	import { Component, Icon } from '@lucide/svelte';

	let {
		icon: ButtonIcon = null,
		buttonLabel,
		dropdownContent
	} = $props<{
		icon?: typeof Component | null;
		buttonLabel: string;
		dropdownContent: Snippet;
	}>();
</script>

<div class="dropdown">
	<div
		tabindex="0"
		role="button"
		class="btn border-base-content/20 align-middle font-normal btn-outline btn-sm hover:bg-primary-content/5"
	>
		{#if ButtonIcon}
			<ButtonIcon class="h-4 w-4" />
		{/if}
		<span class="text-sm">{buttonLabel}</span>
	</div>

	<!--    Dropdown content -->
	<div class="dropdown-content flex-g w-96 shadow-2xl">
		<div class="card mt-1 border border-base-content/15 bg-base-100 shadow-2xl card-xs">
			<div class="card-body p-2">
				<div class="flex flex-col gap-2">
					{@render dropdownContent()}
				</div>
			</div>
		</div>
	</div>
</div>
