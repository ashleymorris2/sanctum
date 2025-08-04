<script lang="ts">
    import {enhance, applyAction} from '$app/forms';
    import {goto} from '$app/navigation';

    let email = $state('');
    let password = $state('');

    const clientErrors = $state({email: '', password: '', error: '', message: ''});

    const errors = $derived.by(() => ({
        email: clientErrors.email,
        password: clientErrors.password,
        error: clientErrors.error,
        message: clientErrors.message
    }));

    const validateEmail = (email: string, cancel: () => void): boolean => {
        if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
            cancel();
            clientErrors.email = 'Please enter a valid email';
            return false;
        }

        clientErrors.email = '';
        clientErrors.message = '';
        return true;
    };

    const onLoginEnhance = ({cancel}) => {

        if (!validateEmail(email, cancel)) {
            return;
        }

        return async ({result, update}) => {
            await update({reset: false, invalidateAll: false});

            if (result.type === 'redirect') {
                // SvelteKit might have redirected or returned success; handle redirection here
                await goto(result.location);
            } else {
                if (result.data.email) clientErrors.email = result.data.email;
                if (result.data.password) clientErrors.password = result.data.password;
                if (result.data.error) clientErrors.error = result.data.error;
                if (result.data.message) clientErrors.message = result.data.message;

                await applyAction(result);
            }
        };
    }
</script>

<div class="w-full max-w-md">
    <div class=" text-center">
        <h1 class="text-4xl font-bold mb-4">Log In</h1>
        <p class="text-base-content/60">Welcome back! Please enter your credentials.</p>
    </div>

    <div class="min-h-[50px] mt-4">
        {#if errors.email || errors.password || errors.error || errors.message}
            <div role="alert" class="alert alert-error">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 shrink-0 stroke-current" fill="none"
                     viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                          d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                <span>{errors.email || errors.password || errors.error || errors.message}</span>
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
                        on:input={() =>{
                             if (validateEmail(email)) {
                                clientErrors.email = '';
                                clientErrors.message = '';
                            }
                        }}
                        required
                        class="input input-lg input-bordered w-full"
                        class:input-error={!!errors.email || !!errors.message}
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
                        on:input={() =>{
                            clientErrors.password = '';
                            clientErrors.message = '';
                        }}
                        required
                        class="input input-lg input-bordered w-full"
                        class:input-error={!!errors.password || errors.message}
                />
            </label>

        </div>
        <button type="submit" class="mt-2 btn btn-lg btn-primary w-full text-primary-content">Continue</button>
    </form>

    <div class="mt-4 text-center">
        <a href="/register" class="link link-hover">Don't have an account? Sign up</a>
    </div>
</div>
