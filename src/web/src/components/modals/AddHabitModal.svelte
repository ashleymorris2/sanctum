<script lang="ts">
    import Modal from './Modal.svelte'
    import type {Habit} from '$lib/habit/types';
    import {CircleCheckBig} from '@lucide/svelte';
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
                <input type="text" bind:value={name} placeholder="Name"
                       class="input w-full text-xl font-semibold border-0 focus:outline-none focus:ring-0 bg-transparent"
                />
            </div>

            <div class="pt-4 pl-2 w-full flex items-center gap-2">
                <div class="dropdown flex-grow">
                    <div tabindex="0" role="button" class="btn align-middle btn-sm btn-outline btn-neutral">
                        <CircleCheckBig class="w-4 h-4"/>
                        <span class="text-sm">{target} {unit} per {frequency === 'daily' ? 'day' : 'week'}</span>
                    </div>
                    <div tabindex="0" class="dropdown-content shadow-2xl">
                        <div class="card bg-base-300 card-xs shadow-2xl">
                            <div class="card-body p-4 ">
                                <div class="flex gap-4">
                                    <div class="flex flex-col w-24">
                                        <label class="label">
                                            <span class="text-sm">Target</span>
                                        </label>
                                        <input type="number"
                                               bind:value={target}
                                               min="1"
                                               placeholder="15"
                                               class="input-simple"
                                        />
                                    </div>
                                    <div class="flex flex-col flex-1">
                                        <label class="label">
                                            <span class="text-sm">Unit</span>
                                        </label>
                                        <input type="text"
                                               bind:value={unit}
                                               placeholder="times"
                                               class="input-simple"
                                        />
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
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
        </div>
    {/snippet}
</Modal>
