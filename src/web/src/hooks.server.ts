import { redirect, type Handle } from '@sveltejs/kit';
import { getRouteAccess } from '$lib/server/auth/routeAccess';
import { verifyAuthToken, refreshAuthToken } from '$lib/server/auth/authTokens';

export const handle: Handle = async ({ event, resolve }) => {
	const authToken = event.cookies.get('auth');
	const refreshToken = event.cookies.get('refresh_token');

	if (getRouteAccess(event.route.id ?? undefined).requiresAuth) {
		if (!authToken && !refreshToken) {
			return redirect(302, '/login');
		}

		if (authToken) {
			const user = await verifyAuthToken(authToken);
			if (user) {
				event.locals.user = user;
				return resolve(event);
			}
		}

		// Try refreshing if access token invalid or missing
		if (refreshToken) {
			const result = await refreshAuthToken(refreshToken);
			if (result?.user && result?.token) {
				event.cookies.set('auth', result.token, {
					httpOnly: true,
					path: '/',
					sameSite: 'lax',
					secure: process.env.NODE_ENV === 'production',
					maxAge: 60 * 15 // 15 min
				});

				event.locals.user = result.user;
				return resolve(event);
			}
		}

		return redirect(302, '/login');
	}

	return resolve(event);
};
