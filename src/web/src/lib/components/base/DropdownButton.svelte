<script lang="ts">
	import type { Snippet } from 'svelte';
	import { Component } from '@lucide/svelte';

	let {
		id = $bindable(''),
		onToggleDropdown,
		icon: ButtonIcon = null,
		buttonLabel,
		dropdownContent,
		styles = 'w-96',
		contentClass = 'flex flex-col gap-2',
		open = $bindable(false)
	} = $props<{
		id?: string;
		onToggleDropdown?: (id: string) => void;
		icon?: typeof Component | null;
		buttonLabel: string;
		dropdownContent: Snippet;
		styles?: string;
		contentClass?: string;
		open?: boolean;
	}>();

	function toggleDropdown(event: MouseEvent) {
		event.preventDefault();
		if (onToggleDropdown) {
			onToggleDropdown(id);
		} else {
			open = !open;
		}
	}
</script>

<details class="dropdown" {open}>
	<summary
		class="btn border-base-content/20 align-middle font-normal btn-outline btn-sm hover:bg-primary-content/5"
		onclick={toggleDropdown}
	>
		{#if ButtonIcon}
			<ButtonIcon class="h-4 w-4" />
		{/if}
		<span class="text-sm">{buttonLabel}</span>
	</summary>

	<!--    Dropdown content -->
	<div class={`dropdown-content flex-g ${styles} shadow-2xl`}>
		<div class="card mt-1 rounded-lg border border-base-content/15 bg-base-100 shadow-2xl card-xs">
			<div class="card-body p-2">
				<div class={contentClass}>
					{@render dropdownContent?.()}
				</div>
			</div>
		</div>
	</div>
</details>
