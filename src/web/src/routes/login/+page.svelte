<script lang="ts">
    import LoginForm from '../../components/LoginForm.svelte';
    import type { ActionData } from './$types';

    export let form: ActionData;
</script>

<div class="min-h-screen bg-base-100 flex ">

    <div class="lg:w-1/2 flex flex-col items-center justify-center p-6">
        <LoginForm {form}/>
    </div>

    <!-- Left panel: only visible on medium and up -->
    <div class="hidden md:flex w-1/2 bg-primary text-primary-content items-center justify-center p-8">
        <div class="space-y-4 text-center">
            <h2 class="text-2xl font-bold">Lorem ipsum dolor sit amet consectetur adipiscing elit.</h2>
            <p class="text-sm"> Sit amet consectetur adipiscing elit quisque faucibus ex.
                Adipiscing elit quisque faucibus ex sapien vitae pellentesque.
            </p>
            <!-- Replace with real image or animation -->
            <img src="/your-image.svg" alt="Illustration" class="mx-auto w-2/3"/>
        </div>
    </div>
</div>


