import type {StorybookConfig} from '@storybook/sveltekit';

const config: StorybookConfig = {
    "stories": [
        "../src/**/*.mdx",
        "../src/**/*.stories.@(js|ts|svelte)"
    ],
    "addons": [
        "@storybook/addon-svelte-csf",
        "@chromatic-com/storybook",
        "@storybook/addon-docs",
        "@storybook/addon-a11y",
        "@storybook/addon-vitest",
				'@storybook/addon-themes',
    ],
    "framework": {
        "name": "@storybook/sveltekit",
        "options": {}
    },
    core: {
        builder: '@storybook/builder-vite',
    },
	async viteFinal(config) {
		// Merge custom configuration into the default config
		const { mergeConfig } = await import('vite');

		return mergeConfig(config, {
			// Add dependencies to pre-optimization
			optimizeDeps: {
				include: ['storybook-dark-mode'],
			},
		});
	},
};
export default config;