import { redirect, type Handle } from '@sveltejs/kit';
import { getRouteAccess } from '$lib/server/auth/routeAccess';
import { verifyAuthToken, refreshAuthToken } from '$lib/server/auth/authTokens';
import { setAuthTokenCookie } from '$lib/server/auth/setCookie';

async function getAuthenticatedUser(authToken: string | undefined) {
	if (!authToken) return null;
	try {
		return await verifyAuthToken(authToken);
	} catch {
		return null;
	}
}

async function tryRefreshUser(event: Parameters<Handle>[0]['event']) {
	const refreshToken = event.cookies.get('refresh_token');
	if (!refreshToken) return null;

	const result = await refreshAuthToken(refreshToken);
	if (!result?.userId || !result?.authToken) return null;

	setAuthTokenCookie(event.cookies, result.authToken);

	return result.userId;
}

export const handle: Handle = async ({ event, resolve }) => {
	const { route } = event;
	const routeAccess = getRouteAccess(route.id ?? undefined);

	if (routeAccess.requiresAuth) {
		const authToken = event.cookies.get('auth_token');
		let user = await getAuthenticatedUser(authToken);

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
