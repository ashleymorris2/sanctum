<script lang="ts">
    export let email = '';
    export let password = '';
    export let onSubmit: (data: { email: string; password: string }) => void = () => {
    };

    function handle(e: Event) {
        e.preventDefault();
        onSubmit({email, password});
    }
</script>

<form on:submit={handle} class="card card-body max-w-sm mx-auto mt-12">
    <h2 class="card-title mb-4 text-xl text-center">Login</h2>

    <label class="label">
        <span class="label-text">Email</span>
    </label>
    <input
            type="email"
            bind:value={email}
            placeholder="you@example.com"
            class="input input-bordered w-full mb-4"
            required
    />

    <label class="label">
        <span class="label-text">Password</span>
    </label>
    <input
            type="password"
            bind:value={password}
            placeholder="••••••"
            class="input input-bordered w-full mb-6"
            required
    />

    <button type="submit" class="btn btn-primary w-full">
        Log In
    </button>
</form>