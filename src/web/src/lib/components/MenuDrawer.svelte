<script lang="ts">
    const {
        links = [],
        pathname = '/',
        title = 'Navigation',
        content = undefined
    } = $props<{
        links?: { label: string; href: string; match?: 'exact' | 'section' }[];
        pathname?: string;
        title?: string;
    }>();
</script>

<div class="drawer lg:drawer-open bg-base-200 ">
    <input id="menu-drawer" type="checkbox" class="drawer-toggle"/>
    <div class="drawer-content flex flex-col">
        {#if content}
            {@render content()}
        {:else}
            <h1 class="text-xl font-bold">{title}</h1>
        {/if}
        <label for="menu-drawer" class="btn btn-primary drawer-button lg:hidden">
            Open drawer
        </label>
    </div>
    <div class="drawer-side">
        <label for="menu-drawer" aria-label="close sidebar" class="drawer-overlay"></label>
        <ul class="menu text-base-content w-64 p-4">
            <li class="text-xl font-bold pb-8">{title}</li>
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