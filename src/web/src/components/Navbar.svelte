<script lang="ts">
    import AddHabitButton from "./AddHabitButton.svelte";

    const {
        title = 'Navigation'
    } = $props<{
        title?: string;
    }>();
</script>

<div class="navbar bg-base-200 shadow-md">
    <div class="navbar-start">
        <p class="text-xl">{title}</p>
    </div>
    <div class="navbar-end">
        <AddHabitButton/>
    </div>
</div>