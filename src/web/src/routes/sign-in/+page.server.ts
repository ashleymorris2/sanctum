import type { Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { AUTH_API_BASE } from '$env/static/private';

export const actions: Actions = {
	default: async ({ request }) => {
		const formData = await request.formData();
		const email = formData.get('email')?.toString().trim();
		const password = formData.get('password')?.toString();
		const formErrors: { [key: string]: string } = {};

		// Validate presence
		if (!email) {
			formErrors.email = 'Email is required';
		}

		if (!password) {
			formErrors.password = 'Password is required';
		}

		if (Object.keys(formErrors).length > 0) {
			return fail(400, { formErrors });
		}

		try {
			const res = await fetch(`${AUTH_API_BASE}/api/sign-in`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ email, password })
			});

			if (!res.ok) {
				const err = await res.json();
				return fail(res.status, { error: err.message || 'Login failed' });
			}

			// Handle session (set cookie, etc.) here

			// Redirect to dashboard
			return redirect(302, '/dashboard');
		} catch (err) {
			console.error('Login error:', err);
			return fail(500, { error: 'Server error during login' });
		}
	}
};
