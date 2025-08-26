<script lang="ts">
    import {Plus} from '@lucide/svelte';
    import AddHabitModal from "./AddHabitModal.svelte";
    import type {Habit} from "$lib/habit/types";

    let open = $state(false);

    function onConfirm(habit: Habit) {
        alert(`submitted ${habit.name}`);
    }
</script>

<button
        class="btn lg:btn-sm btn-primary text-white"
        on:click={() => (open = true)}
>
    <Plus size="16"/>
    Add habit
</button>

<AddHabitModal bind:open={open} onConfirm={onConfirm}/>