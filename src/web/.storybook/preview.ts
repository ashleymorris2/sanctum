import type { Preview, SvelteRenderer } from '@storybook/sveltekit';
import { withThemeByDataAttribute } from '@storybook/addon-themes';
import '../src/app.css';

const preview: Preview = {
	parameters: {
		controls: {
			matchers: {
				color: /(background|color)$/i,
				date: /Date$/i
			}
		}
	},
	decorators: [
		withThemeByDataAttribute<SvelteRenderer>({
			themes: {
				light: 'sanctum',
				dark: 'dark',
				sanctum: 'sanctum'
			},
			defaultTheme: 'light'
			// attributeName: 'data-theme',
		})
	]
};

export default preview;
