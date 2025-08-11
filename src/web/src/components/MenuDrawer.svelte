<script lang="ts">
    export let links: { label: string; href: string }[] = []; // Links to show in the MenuDrawer
    export let pathname: string = '/'; // The current pages pathname
    export let title: string = 'Navigation';
</script>

<div class="drawer lg:drawer-open">
    <input id="my-drawer-2" type="checkbox" class="drawer-toggle"/>
    <div class="drawer-content flex flex-col items-center justify-center">
        <!-- Page content here -->
        <label for="my-drawer-2" class="btn btn-primary drawer-button lg:hidden">
            Open drawer
        </label>
    </div>
    <div class="drawer-side">
        <label for="my-drawer-2" aria-label="close sidebar" class="drawer-overlay"></label>
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