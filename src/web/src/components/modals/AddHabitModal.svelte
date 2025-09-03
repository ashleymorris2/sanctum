<script lang="ts">
    import Modal from './Modal.svelte'
    import type {Habit} from '$lib/habit/types';
    import {validateHabit} from "$lib/habit/validation";
    import TargetInputButton from "../button/TargetInputButton.svelte";
    import HabitTypeButton from "../button/HabitTypeButton.svelte";

    let {
        open = $bindable(false),
        habit = null as Habit | null,
        onConfirm = function () {
        },
    } = $props();

    let name = $state('');
    let description = $state('');
    let target = $state('');
    let unit = $state('');
    let frequency = $state('');

    function resetFrom(h: Habit | null) {
        name = h?.name ?? '';
        description = h?.description ?? '';
        target = h?.target?.toString() ?? '1';
        unit = h?.unit ?? 'time';
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

<Modal bind:open bind:isValid={habitIsValid} onConfirm={confirm}>
    {#snippet content()}
        <div class="p-4">
            <div class="w-full ">
                <input type="text"
                       bind:value={name}
                       placeholder="Name"
                       class="w-full text-xl font-semibold bg-transparent outline-none focus:ring-0 pl-2 pt-2 pb-2"
                />
                <input type="text"
                       bind:value={description}
                       placeholder="Description"
                       class="w-full txt-sm bg-transparent outline-none focus:ring-0 pl-2 pb-6"
                />
            </div>

            <div class="flex items-center pt-4 gap-4">
                <HabitTypeButton/>
                <div>
                    <TargetInputButton bind:frequency={frequency} bind:target={target} bind:unit={unit}/>
                </div>
            </div>


        </div>
    {/snippet}
</Modal>
