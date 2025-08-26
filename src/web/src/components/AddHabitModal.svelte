<script lang="ts">
    import Modal from './Modal.svelte'
    import type {Habit} from '$lib/habit/types';
    import {validateHabit} from "$lib/habit/validation";

    let {
        open = $bindable(false),
        habit = null as Habit | null,
        onConfirm = function () {
        },
    } = $props();

    let name = $state('');
    let target = $state('');
    let unit = $state('');
    let frequency = $state('');

    $effect(() => {
        name = habit?.name ?? '';
        target = habit?.target?.toString() ?? '';
        unit = habit?.unit ?? '';
        frequency = habit?.frequency ?? '';
    });

    const input = $derived.by(() => ({
        id: habit?.id,
        name,
        target,
        unit,
        frequency
    }));

    let habitIsValid = $derived.by(() => {
        return validateHabit(input).isValid;
    });

    function confirm() {
        if (!habitIsValid) return;
        onConfirm?.(input);
        open = false;
    }
</script>

<Modal title="Add Habit" bind:open bind:isValid={habitIsValid} onConfirm={confirm}>
    {#snippet content()}
        <div class="bg-base-200 p-6">
            <div class="w-full flex items-center gap-2">
                <label class="floating-label flex-grow">
                    <span>Name</span>
                    <input type="text" bind:value={name} placeholder="Name" class="input-simple w-full"
                           aria-invalid="true"/>
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
                    <input type="text" bind:value={target} placeholder="15" class="input-simple w-16"/>
                </label>
                <label class="floating-label flex-grow">
                    <span>Unit</span>
                    <input type="text" bind:value={unit} placeholder="reps" class="input-simple w-full"/>
                </label>
                <label class="floating-label flex-grow">
                    <span>Frequency</span>
                    <input type="text" bind:value={frequency} placeholder="daily" class="input-simple w-full"/>
                </label>
                <label class="floating-label flex-grow">
                    <span>Frequency</span>
                    <input type="text" placeholder="average" class="input-simple w-full"/>
                </label>
            </div>
        </div>
    {/snippet}
</Modal>
