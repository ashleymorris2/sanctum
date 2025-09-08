<script lang="ts">
	let {
		open = $bindable(false),
		content = undefined,
		isValid = $bindable(false),
		onConfirm = function () {}
	} = $props();

	let modal: HTMLDialogElement | null = null;

	$effect(() => {
		if (!modal) return;
		if (open && !modal.open) modal.showModal();
		if (!open && modal.open) modal.close();
	});

	function onClose() {
		open = false;
	}

	const close = () => (open = false);
</script>

<dialog bind:this={modal} class="modal" onclose={onClose}>
	<div class="modal-box overflow-visible rounded-xl p-0 shadow-xl">
		{@render content?.()}

		<div class="divider m-0 h-0"></div>
		<div class="modal-action mt-0 p-4">
			<div class="flex items-center gap-2">
				<button class="btn btn-sm btn-soft" onclick={close}>Cancel</button>
				<button class="btn btn-primary btn-sm" onclick={onConfirm} disabled={!isValid}
					>Add Habit</button
				>
			</div>
		</div>
	</div>
</dialog>
