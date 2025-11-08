import { redirect, type Handle } from '@sveltejs/kit';
import { getRouteAccess } from '$lib/server/auth/routeAccess';
import { verifyAccessToken, refreshAccessToken } from '$lib/server/auth/accessTokens';
import { setAuthTokenCookie } from '$lib/server/auth/setCookie';

async function getAuthenticatedUser(accessToken: string | undefined) {
	if (!accessToken) return null;
	try {
		return await verifyAccessToken(accessToken);
	} catch {
		return null;
	}
}

async function tryRefreshUser(event: Parameters<Handle>[0]['event']) {
	const refreshToken = event.cookies.get('refresh_token');
	if (!refreshToken) return null;

	const result = await refreshAccessToken(refreshToken);
	if (!result?.userId || !result?.authToken) return null;

	setAuthTokenCookie(event.cookies, result.authToken);

	return result.userId;
}

export const handle: Handle = async ({ event, resolve }) => {
	const { route } = event;
	const routeAccess = getRouteAccess(route.id ?? undefined);

	if (routeAccess.requiresAuth) {
		const accessToken = event.cookies.get('auth_token');
		let user = await getAuthenticatedUser(accessToken);

		if (!user) {
			user = await tryRefreshUser(event);
		}

		if (!user) {
			return redirect(302, '/login');
		}

		event.locals.user = user;
	}

	return resolve(event);
};
