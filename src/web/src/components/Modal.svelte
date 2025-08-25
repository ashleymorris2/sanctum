<script lang="ts">
    type Props = {
        open?: boolean;
        title?: string;
        content?: () => Snippet;
    };

    let {
        open = false,
        isValid = true,
        title = "",
        content = undefined
    } = $props<Props>();

    let modal: HTMLDialogElement | null = null;

    $effect(() => {
        if (!modal) return;
        if (open && !modal.open) modal.showModal();
        if (!open && modal.open) modal.close();
    });
</script>

<dialog bind:this={modal} class="modal" on:close={() => (open = false)}>
    <div class="modal-box p-0">
        <h3 class="text-lg font-bold p-6">{title}</h3>
        <div class="divider m-0 h-0"/>

        {@render content?.()}

        <div class="divider m-0 h-0"/>
        <div class="modal-action mt-0 p-6">
            <div class="flex items-center gap-2">
                <button class="btn btn-ghost" on:click={close}>Cancel</button>
                <button class="btn btn-primary" on:click={close} disabled={!isValid}>Save</button>
            </div>
        </div>
    </div>
</dialog>