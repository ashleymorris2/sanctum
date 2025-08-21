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
    <div class="modal-box p-0">
        <h3 class="text-lg font-bold p-6">Create a habit</h3>
        <div class="divider m-0 h-0"/>
        <div class="bg-base-200 p-6">
            <div class="w-full flex items-center gap-2">
                <label class="floating-label flex-grow">
                    <span>Name</span>
                    <input type="text" placeholder="Name" class="input w-full"/>
                </label>
                <button class="btn btn-square btn-primary btn-outline shrink-0">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24" stroke-width="2.5"
                         stroke="currentColor" class="size-[1.2em]">
                        <path stroke-linecap="round" stroke-linejoin="round"
                              d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z"/>
                    </svg>
                </button>
            </div>

            <div class="pt-4 w-full flex items-center gap-2">
                <label class="floating-label flex-grow">
                    <span>Target</span>
                    <input type="text" placeholder="15" class="input-floating w-16"/>
                </label>
                <label class="floating-label flex-grow">
                    <span>Unit</span>
                    <input type="text" placeholder="reps" class="input w-full"/>
                </label>
                <label class="floating-label flex-grow">
                    <span>Frequency</span>
                    <input type="text" placeholder="daily" class="input w-full"/>
                </label>
                <label class="floating-label flex-grow">
                    <span>Frequency</span>
                    <input type="text" placeholder="average" class="input w-full"/>
                </label>
            </div>
        </div>

        <div class="divider m-0 h-0"/>

        <div class="modal-action mt-0 p-6">
            <div class="flex items-center gap-2">
                <button class="btn btn-ghost" on:click={close}>Cancel</button>
                <button class="btn btn-primary" on:click={close} disabled={true}>
                    Save
                </button>
            </div>
        </div>
    </div>
</dialog>