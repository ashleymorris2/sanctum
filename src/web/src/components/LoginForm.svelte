<script lang="ts">
    import {deserialize} from '$app/forms';

    let emailError = '';
    let passwordError = '';
    let globalError = '';

    async function handleSubmit(event: Event) {
        event.preventDefault();
        const formData = new FormData(event.target as HTMLFormElement);
        const res = await fetch('/login', {
                method: 'POST',
                body: formData
            }
        );

        // Handle redirect manually
        if (res.redirected) {
            window.location.href = res.url;
            return;
        }

        const result = deserialize(await res.text());
        const errorData = result.data;

        // Reset previous errors
        emailError = '';
        passwordError = '';
        globalError = '';

        // Assign errors from backend
        if (typeof errorData === 'object' && errorData !== null) {
            emailError = errorData.email ?? '';
            passwordError = errorData.password ?? '';
            globalError = errorData.globalError ?? '';
        } else if (typeof errorData === 'string') {
            globalError = errorData;
        }
    }
</script>

<div class="w-full max-w-md">
    <div class=" text-center">
        <h1 class="text-4xl font-bold mb-4">Log In</h1>
        <p class="text-base-content/60">Welcome back! Please enter your credentials.</p>
    </div>

    <div class="min-h-[50px] mt-4">
        {#if emailError || passwordError || globalError}
            <div role="alert" class="alert alert-error">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 shrink-0 stroke-current" fill="none"
                     viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                          d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                <span>{emailError || passwordError || globalError}</span>
            </div>
        {/if}
    </div>

    <form method="POST" on:submit={handleSubmit} class="card mx-auto mt-2 card-body max-w-sm" novalidate>
        <div>
            <label class="floating-label mb-4">
                <span class="label">
                    <span class="label-text">Email</span>
                </span>
                <input
                        id="email"
                        type="email"
                        name="email"
                        placeholder="Email address"
                        autocomplete="email"
                        required
                        class="input input-lg  input-bordered w-full"
                        class:input-error={ !!emailError || !!globalError}
                />
            </label>
            <label class="floating-label mt-4">
                <span class="label">
                    <span class="label-text">Password</span>
                </span>
                <input
                        type="password"
                        name="password"
                        placeholder="Password"
                        autocomplete="current-password"
                        required
                        class="input input-lg input-bordered w-full input-error"
                        class:input-error={ !!passwordError || !!globalError}
                />
            </label>
            <!--            <p class="text-sm">{true || '\u00A0'}</p>-->
        </div>
        <button type="submit" class="mt-2 btn btn-lg btn-primary w-full text-primary-content">Continue</button>
    </form>

    <div class="mt-4 text-center">
        <a href="/register" class="link link-hover">Don't have an account? Sign up</a>
    </div>
</div>
