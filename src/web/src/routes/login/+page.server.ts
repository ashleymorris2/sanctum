import type { Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { AUTH_API_BASE } from '$env/static/private';

export const actions: Actions = {
	default: async ({ request, cookies }) => {
		const formData = await request.formData();
		const email = formData.get('email')?.toString().trim();
		const password = formData.get('password')?.toString();

		if (!email && !password) {
			return fail(400, { error: 'Email and Password is required' });
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

			const data = await res.json();

			if (!res.ok) {
				return fail(res.status, data);
			}

			const { authToken, refreshToken, refreshTokenTTL } = data;

			cookies.set('refresh_token', refreshToken, {
				path: '/',
				httpOnly: true,
				sameSite: 'lax',
				secure: process.env.NODE_ENV === 'production',
				maxAge: refreshTokenTTL // 7 days
			});

			return {
				success: true,
				authToken,
				location: '/dashboard'
			};
		} catch {
			return fail(500, { error: 'Server error during login' });
		}
	}
};
