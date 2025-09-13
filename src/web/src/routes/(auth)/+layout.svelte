<script lang="ts">
    import { page } from '$app/state';
    import Navbar from "../../lib/components/Navbar.svelte";
    import MenuDrawer from "../../lib/components/MenuDrawer.svelte";

    let {children} = $props();
    let pathname = $derived(page.url.pathname);
    let user = null;
    let drawerOpen = false;
    const links = [
        {label: 'Dashboard', href: '/dashboard'},
        {label: 'Habits', href: '/habits'},
        {label: 'Settings', href: '/settings'}
    ];
    let currentLinkLabel = $derived(links.find(link =>
        pathname === link.href || (link.href !== '/' && pathname.startsWith(link.href))
    )?.label ?? 'My App');
</script>

<div>
    <MenuDrawer {links} {pathname} title="Sanctum" bind:open={drawerOpen}>
        {#snippet content()}
        <div>
            <Navbar title={currentLinkLabel} {user} on:menu={() => (drawerOpen = true)}/>
            <main>
                {@render children?.()}
            </main>
        </div>
        {/snippet}
    </MenuDrawer>
</div>