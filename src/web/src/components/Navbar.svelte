<script lang="ts">
    const {
        title = 'Navigation'
    } = $props<{
        title?: string;
    }>();
</script>

<div class="navbar bg-base-200 shadow-md">
    <a class="btn btn-ghost text-xl">{title}</a>
</div>