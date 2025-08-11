<script lang="ts">
    export let links: { label: string; href: string }[] = []; // Links to show in the MenuDrawer
    export let pathname: string = '/'; // The current pages pathname
    export let title: string = 'Navigation';

    let { children } = $props();
</script>

<div class="drawer lg:drawer-open">
    <input id="menu-drawer" type="checkbox" class="drawer-toggle"/>
    <div class="drawer-content flex flex-col items-center justify-center">
        {#if children}
            {@render children()}
        {:else}
            <span>No content found</span>
        {/if}
        <label for="menu-drawer" class="btn btn-primary drawer-button lg:hidden">
            Open drawer
        </label>
    </div>
    <div class="drawer-side">
        <label for="menu-drawer" aria-label="close sidebar" class="drawer-overlay"></label>
        <ul class="menu bg-base-200 text-base-content min-h-full w-80 p-4">
            {#each links as link (link.href)}
                <li>
                    <a href={link.href} class:active-item={pathname.startsWith(link.href)}>
                        {link.label}
                    </a>
                </li>
            {/each}
        </ul>
    </div>
</div>