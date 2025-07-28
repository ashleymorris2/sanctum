<script lang="ts">
    import LoginForm from '../../components/LoginForm.svelte';
</script>

<div class="min-h-screen bg-base-100 flex ">

    <div class="lg:w-1/2 flex flex-col items-center justify-center p-6">
        <div class="w-full max-w-md">
            <div class="text-center mb-8">
                <h1 class="text-4xl font-bold">Login</h1>
                <p class="text-base-content/60">Welcome back! Please login to your account.</p>
            </div>

            <LoginForm/>

            <div class="text-center mt-6">
                <a href="/register" class="link link-hover">Don't have an account? Sign up</a>
            </div>
        </div>
    </div>

    <!-- Left panel: only visible on medium and up -->
    <div class="hidden md:flex w-1/2 bg-primary text-white items-center justify-center p-8">
        <div class="space-y-4 text-center">
            <h2 class="text-2xl font-bold">New Scheduling And Routing Options</h2>
            <p class="text-sm">We also updated the format of podcasts and rewards.</p>
            <!-- Replace with real image or animation -->
            <img src="/your-image.svg" alt="Illustration" class="mx-auto w-2/3"/>
        </div>
    </div>
</div>


