<script lang="ts">
    import { enhance } from '$app/forms';
    import { goto } from '$app/navigation';

    let email = $state('');
    let password = $state('');

    const clientErrors = $state({email: '', password: ''});
    let serverErrors = $state({email: '', password: '',});

    // const formData = $derived(page.form as PageData['form']);

    const errors = $derived.by(() => ({
        email: clientErrors.email ?? serverErrors?.email,
        password: clientErrors.password ?? serverErrors?.password,
        globalError: clientErrors.globalError ?? serverErrors?.globalError
    }));

    const onLoginEnhance = ({cancel}) => {

        if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
            cancel();
            clientErrors.email = 'Please enter a valid email';
            return;
        }

        return async ({result, update}) => {
            // // Client-side validation before server
            clientErrors.email = '';
            clientErrors.password = '';
            serverErrors.email = '';
            serverErrors.password = '';


            console.log('form before update', result);

            // Let SvelteKit populate `form` prop and handle `fail()` or `redirect()`
            await update({reset: false, invalidateAll: false});

            console.log('form after update', result);

            if (result.type === 'redirect') {
                // SvelteKit might have redirected or returned success; handle redirection here
                await goto(result.location);
            } else if (result.type === 'failure') {
                // At this point, `form` contains your server-side error object
                // e.g., `form.email`, `form.password`, `form.globalError`
                // you can also set focus or tools here
                if (result.data.email) serverErrors.email = result.data.email;
                if (result.data.password) serverErrors.password = result.data.password;
                console.log(result);
            } else {
                // successful submission without redirect
            }
        };
    };
</script>

<div class="w-full max-w-md">
    <div class=" text-center">
        <h1 class="text-4xl font-bold mb-4">Log In</h1>
        <p class="text-base-content/60">Welcome back! Please enter your credentials.</p>
    </div>

    <div class="min-h-[50px] mt-4">
        {#if errors.email || errors.password || errors.globalError}
            <div role="alert" class="alert alert-error">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 shrink-0 stroke-current" fill="none"
                     viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                          d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                <span>{errors.email || errors.password || errors.globalError}</span>
            </div>
        {/if}
    </div>

    <form method="POST" action="/login" use:enhance={onLoginEnhance} class="card mx-auto mt-2 card-body max-w-sm"
          novalidate>
        <div>
            <label class="floating-label mb-4">
                <span class="label">
                    <span class="label-text">Email</span>
                </span>
                <input
                        id="email"
                        type="email"
                        name="email"
                        bind:value={email}
                        placeholder="Email address"
                        autocomplete="email"
                        required
                        class="input input-lg input-bordered w-full"
                        class:input-error={!!errors.email}
                />
            </label>
            <label class="floating-label mt-4">
                <span class="label">
                    <span class="label-text">Password</span>
                </span>
                <input
                        type="password"
                        name="password"
                        bind:value={password}
                        placeholder="Password"
                        autocomplete="current-password"
                        required
                        class="input input-lg input-bordered w-full"
                        class:input-error={!!errors.password}
                />
            </label>

        </div>
        <button type="submit" class="mt-2 btn btn-lg btn-primary w-full text-primary-content">Continue</button>
    </form>

    <div class="mt-4 text-center">
        <a href="/register" class="link link-hover">Don't have an account? Sign up</a>
    </div>
</div>
