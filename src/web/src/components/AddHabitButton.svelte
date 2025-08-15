<script lang="ts">
    import {Plus} from '@lucide/svelte';
    import NewHabitModal from "./NewHabitModal.svelte";

    let open = $state(false);
</script>

<button
        class="btn lg:btn-sm btn-primary text-primary-content"
        on:click={() => (open = true)}
>
    <Plus size="16"/>
    Add habit
</button>

<NewHabitModal bind:open={open}/>