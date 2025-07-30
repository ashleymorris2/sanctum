<script module>
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import SignInPage from '../routes/login/+page.svelte';

    const { Story } = defineMeta({
        title: 'Pages/login',
        component: SignInPage,
        parameters: {
            layout: 'fullscreen',
        },
    });
</script>

<Story name="Default" args={{}}>
    {#snippet children(args)}
        <SignInPage {...args} />
    {/snippet}
</Story>

