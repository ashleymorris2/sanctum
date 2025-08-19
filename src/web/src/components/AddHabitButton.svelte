<script lang="ts">
    import {Plus} from '@lucide/svelte';
    import AddHabitModal from "./AddHabitModal.svelte";

    let open = $state(false);
</script>

<button
        class="btn lg:btn-sm btn-primary text-white"
        on:click={() => (open = true)}
>
    <Plus size="16"/>
    Add habit
</button>

<AddHabitModal bind:open={open}/>