<script lang="ts">
    import AddHabitButton from "./button/AddHabitButton.svelte";

    const {
        title = 'Navigation'
    } = $props<{
        title?: string;
    }>();
</script>

<div class="navbar bg-base-200 relative z-10">
    <div class="navbar-start">
        <p class="text-xl">{title}</p>
    </div>
    <div class="navbar-end">
        <AddHabitButton/>
    </div>
</div>