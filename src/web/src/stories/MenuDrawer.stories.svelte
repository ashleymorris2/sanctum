<script module lang="ts">
    import {defineMeta} from '@storybook/addon-svelte-csf';
    import MenuDrawer from '$lib/components/MenuDrawer.svelte';

    const {Story} = defineMeta({
        title: 'Navigation/SideMenu',
        component: MenuDrawer,
        parameters: {
            layout: 'fullscreen',
        },
        args: {
            title: 'Sanctum',
            links: [
                {label: 'Habits', href: '/habits'},
                {label: 'Metrics', href: '/metrics'}
            ]
        },
    });
</script>

<Story name="HabitsActive" args={{ pathname: '/habits' }} let:args>
    <MenuDrawer {...args}/>
</Story>