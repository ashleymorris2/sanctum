<script lang="ts">
    import Modal from './Modal.svelte'
    import type {Habit} from '$lib/habit/types';
    import {validateHabit} from "$lib/habit/validation";
    import TargetInputButton from "../button/TargetInputButton.svelte";

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
                       bind:value={name}
                       placeholder="Description"
                       class="w-full txt-sm font-semibold bg-transparent outline-none focus:ring-0 pl-2 pb-6"
                />
            </div>

            <div class="pt-4 pl-2 w-full flex items-center gap-2">
                <TargetInputButton bind:frequency={frequency} bind:target={target} bind:unit={unit}/>
            </div>
            <!--                <label class="floating-label flex-grow">-->
            <!--                    <span>Unit</span>-->
            <!--                    <input type="text" bind:value={unit} placeholder="reps" class="input-simple w-full"/>-->
            <!--                </label>-->
            <!--                <label class="floating-label flex-grow">-->
            <!--                    <span>Frequency</span>-->
            <!--                    <input type="text" bind:value={frequency} placeholder="daily" class="input-simple w-full"/>-->
            <!--                </label>-->
            <!--                <label class="floating-label flex-grow">-->
            <!--                    <span>Frequency</span>-->
            <!--                    <input type="text" placeholder="average" class="input-simple w-full"/>-->
            <!--                </label>-->

        </div>
    {/snippet}
</Modal>
