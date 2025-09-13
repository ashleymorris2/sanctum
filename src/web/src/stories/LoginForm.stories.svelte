<script module lang="ts">
    import { defineMeta } from '@storybook/addon-svelte-csf';
    import LoginForm from '$lib/components/LoginForm.svelte';

    const { Story } = defineMeta({
        title: 'Auth/LoginForm',
        component: LoginForm
    });
</script>

<Story name="Default" let:args>
    <LoginForm
            email={args.email}
            password={args.password}
            onSubmit={args.onSubmit}
    />
</Story>