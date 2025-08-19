<script lang="ts">
    const title = 'Placeholder Page';
    const description = 'This is a placeholder page for development purposes.';
</script>

<div class="min-h-screen bg-base  pt-4">
    <div class="max-w-4xl mx-auto">
        <div class="bg-white rounded-lg shadow-md p-8">
            <h1 class="text-3xl font-bold text-gray-900 mb-4">{title}</h1>
            <p class="text-gray-600 mb-6">{description}</p>

            <div class="space-y-4">
                <!-- Placeholder content blocks -->
                <div class="h-8 bg-gray-200 rounded animate-pulse"></div>
                <div class="h-32 bg-gray-200 rounded animate-pulse"></div>
                <div class="h-8 bg-gray-200 rounded animate-pulse"></div>
            </div>

            <div class="mt-8">
                <button
                        class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded transition-colors"
                        type="button"
                >
                    Sample Button
                </button>
            </div>
        </div>
    </div>
</div>