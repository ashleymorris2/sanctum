<script lang="ts">
    import {enhance} from '$app/forms';

    export let form: {
        errorMessage?: string;
    };
</script>

<div class="w-full max-w-md">
    <div class=" text-center">
        <h1 class="text-4xl font-bold mb-8">Log In</h1>
        <p class="text-base-content/60">Welcome back! Please enter your credentials.</p>
    </div>

    <form method="POST" use:enhance class="card mx-auto mt-8 card-body max-w-sm" novalidate>
        <div>
            <label class="floating-label mb-4">
                <span class="label">
                    <span class="label-text">Email</span>
                </span>
                <input
                        id="email"
                        class="input input-lg input-bordered w-full {form?.errorMessage ? 'input-error' : ''}"
                        type="email"
                        name="email"
                        placeholder="Email address"
                        autocomplete="email"
                        required
                />
            </label>
            <label class="floating-label mt-4">
                <span class="label">
                    <span class="label-text">Password</span>
                </span>
                <input
                        class="input input-lg input-bordered w-full {form?.errorMessage ? 'input-error' : ''}"
                        type="password"
                        name="password"
                        placeholder="Password"
                        autocomplete="current-password"
                        required
                />
            </label>
            <p class="text-sm">{form?.errorMessage || '\u00A0'}</p>
        </div>
        <button type="submit" class="mt-2 btn btn-lg btn-primary w-full text-primary-content">Continue</button>
    </form>

    <div class="mt-4 text-center">
        <a href="/register" class="link link-hover">Don't have an account? Sign up</a>
    </div>
</div>
