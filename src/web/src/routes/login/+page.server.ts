import type { Actions } from './$types';
import { fail } from '@sveltejs/kit';
import { env } from '$env/dynamic/private';
import { setRefreshTokenCookie } from '$lib/server/auth/setCookie';

export const actions: Actions = {
	default: async ({ request, cookies }) => {
		const formData = await request.formData();
		const email = formData.get('email')?.toString().trim();
		const password = formData.get('password')?.toString();

		if (!email && !password) {
			return fail(400, { error: 'An Email and Password are required' });
		}
		if (!email) {
			return fail(400, { email: 'An email is required' });
		}
		if (!password) {
			return fail(400, { password: 'A password is required' });
		}

		try {
			const res = await fetch(`${env.AUTH_API_BASE}/api/login`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ email, password })
			});

			const data = await res.json();
			if (!res.ok) {
				return fail(res.status, data);
			}

			const { authToken, refreshToken, refreshTokenTTL } = data;

			setRefreshTokenCookie(cookies, refreshToken, refreshTokenTTL);

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
