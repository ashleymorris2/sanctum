<script module lang="ts">
    import {defineMeta} from '@storybook/addon-svelte-csf';
    import MenuDrawer from '../components/MenuDrawer.svelte';

    const {Story} = defineMeta({
        title: 'Navigation/MenuDrawer',
        component: MenuDrawer,
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

<Story name="NoActive" args={{ pathname: '/' }} let:args>
    <MenuDrawer {...args}/>
</Story>