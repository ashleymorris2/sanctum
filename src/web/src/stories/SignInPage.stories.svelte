<script module>
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import SignInPage from '../routes/sign-in/+page.svelte';

    const { Story } = defineMeta({
        title: 'Pages/sign-in',
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

