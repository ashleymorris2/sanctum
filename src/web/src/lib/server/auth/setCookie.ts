import type { Cookies } from '@sveltejs/kit';

export function setAuthTokenCookie(cookies: Cookies, token: string) {
	cookies.set('auth_token', token, {
		httpOnly: true,
		secure: process.env.NODE_ENV === 'production',
		sameSite: 'lax',
		path: '/',
		maxAge: 60 * 15 // 15 minutes
	});
}

export function setRefreshTokenCookie(cookies: Cookies, token: string, ttl: number) {
	cookies.set('refresh_token', token, {
		httpOnly: true,
		secure: process.env.NODE_ENV === 'production',
		sameSite: 'lax',
		path: '/',
		maxAge: ttl
	});
}
