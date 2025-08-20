<script lang="ts">

    // runes: a bindable prop for open/close
    let {open = $bindable(false)} = $props();

    let modal: HTMLDialogElement | null = null;

    $effect(() => {
        if (!modal) return;
        if (open && !modal.open) modal.showModal();
        if (!open && modal.open) modal.close();
    });

    function onClose() {
        open = false; // updates parent via bind:
    }
</script>

<dialog bind:this={modal} class="modal" on:close={onClose}>
    <div class="modal-box p-0 shadow-xl rounded-1 ">
        <div class="p-2">
            <input
                    placeholder="Habit name"
                    class="input w-full text-lg font-semibold border-0 focus:outline-none focus:ring-0 bg-transparent"
            />

        </div>

        <div class="divider"></div>

        <div class="modal-action p-2">
            <div class="flex items-center gap-2">
                <button class="btn btn-ghost" on:click={close}>Cancel</button>
                <button class="btn btn-primary" on:click={close} disabled={true}>
                    Add task
                </button>
            </div>
        </div>
    </div>
</dialog>