import type { Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { AUTH_API_BASE } from '$env/static/private';

export const actions: Actions = {
	default: async ({ request, cookies }) => {
		const formData = await request.formData();
		const email = formData.get('email')?.toString().trim();
		const password = formData.get('password')?.toString();

		if (!email && !password) {
			return fail(400, { globalError: 'Email and Password is required' });
		}
		if (!email) {
			return fail(400, { email: 'Email is required' });
		}
		if (!password) {
			return fail(400, { password: 'Password is required' });
		}

		try {
			const res = await fetch(`${AUTH_API_BASE}/api/login`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ email, password })
			});

			if (!res.ok) {
				const errorData = await res.json();
				return fail(res.status, errorData);
			}

			const { token } = await res.json();

			cookies.set('auth', token, {
				path: '/',
				httpOnly: true,
				sameSite: 'lax',
				secure: process.env.NODE_ENV === 'production',
				maxAge: 60 * 60 * 24 * 7 // 7 days
			});
		} catch (err) {
			console.error('Login error:', err);
			return fail(500, { error: 'Server error during login' });
		}

		// Redirect to dashboard
		return redirect(302, '/dashboard');
	}
};
