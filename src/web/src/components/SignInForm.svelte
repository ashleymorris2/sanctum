<script lang="ts">
    import { enhance } from '$app/forms';
</script>

<div class="w-full max-w-md">
    <div class="mb-8 text-center">
        <h1 class="text-4xl font-bold">Sign in</h1>
        <p class="text-base-content/60">Welcome back! Please enter your credentials.</p>
    </div>

    <form method="POST" use:enhance action="/sign-in" class="card mx-auto mt-12 card-body max-w-sm">

        <label class="label" for="email">
            <span class="label-text">Email</span>
        </label>
        <input
                type="email"
                name="email"
                placeholder="you@example.com"
                class="input-bordered input mb-4 w-full"
                required
        />

        <label class="label" for="password">
            <span class="label-text">Password</span>
        </label>
        <input
                type="password"
                name="password"
                placeholder="••••••"
                class="input-bordered input mb-6 w-full"
                required
        />

        <button type="submit" class="btn w-full btn-primary"> Sign in</button>
    </form>

    <div class="mt-6 text-center">
        <a href="/register" class="link link-hover">Don't have an account? Sign up</a>
    </div>
</div>
