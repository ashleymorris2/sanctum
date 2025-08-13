<script lang="ts">
    import {Plus} from '@lucide/svelte';

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
        <button class="btn lg:btn-sm btn-primary text-primary-content">
            <Plus size="16"/>
            Add habit
        </button>
    </div>
</div>