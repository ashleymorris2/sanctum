<script module lang="ts">
    import {defineMeta} from '@storybook/addon-svelte-csf';
    import Navbar from '../components/Navbar.svelte';

    const {Story} = defineMeta({
        title: 'Navigation/Navbar',
        component: Navbar,
        parameters: {
            layout: 'fullscreen',
        },
        args: {
            title: 'Sanctum'
        },
    });
</script>

<Story name="Navbar" let:args>
    <Navbar {...args}/>
</Story>
