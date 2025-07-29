<script module lang="ts">
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import SignInForm from '../components/SignInForm.svelte';

    const { Story } = defineMeta({
        title: 'Auth/SignInForm',
        component: SignInForm
    });
</script>

<Story name="Default" let:args>
    <SignInForm
            email={args.email}
            password={args.password}
            onSubmit={args.onSubmit}
    />
</Story>