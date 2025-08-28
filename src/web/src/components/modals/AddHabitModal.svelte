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

    function resetFrom(h: Habit | null) {
        name = h?.name ?? '';
        target = h?.target?.toString() ?? '';
        unit = h?.unit ?? '';
        frequency = h?.frequency ?? 'daily';
    }

    $effect(() => {
        if (open) resetFrom(habit);
    });

    const input = $derived.by(() => ({
        id: habit?.id,
        name,
        target: target === '' ? null : Number(target),
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
            </div>

            <div class="pt-4 w-full flex items-center gap-2">
                <label class="floating-label flex-grow">
                    <span>Target</span>
                    <input type="number" bind:value={target} min="1" placeholder="15" class="input-simple w-16"/>
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
