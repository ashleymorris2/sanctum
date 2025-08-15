<script lang="ts">

    // runes: a bindable prop for open/close
    let { open = $bindable(false) } = $props();

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
    <div class="modal-box">
        <h3 class="text-lg font-bold">Create a habit</h3>
        <div class="pt-2 w-full">
            <label class="floating-label">
                <span>Name</span>
                <input type="text" placeholder="Name" class="input input-md"/>
            </label>
        </div>
        <div class="modal-action">
            <form method="dialog">
                <button class="btn">Close</button>
            </form>
        </div>
    </div>
</dialog>