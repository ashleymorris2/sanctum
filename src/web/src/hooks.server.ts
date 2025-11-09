import { redirect, type Handle } from '@sveltejs/kit';
import { getRouteAccess } from '$lib/server/auth/routeAccess';
import { refreshAccessToken, verifyAccessTokenLocal } from '$lib/server/auth/accessTokens';
import { setAccessTokenCookie as setAccessTokenCookie } from '$lib/server/auth/setCookie';
import type { User } from '$lib/types/user';

async function getAuthenticatedUser(accessToken: string | undefined): Promise<User | null> {
	if (!accessToken) return null;
	const claims = await verifyAccessTokenLocal(accessToken); // returns claims | null
	if (!claims) return null;

	return {
		id: claims.sub as string
	};
}

async function tryRefreshUser(event: Parameters<Handle>[0]['event']): Promise<User | null> {
	const refreshToken = event.cookies.get('refresh_token');
	if (!refreshToken) return null;

	const accessToken = await refreshAccessToken(refreshToken);
	if (!accessToken) return null;

	const claims = await verifyAccessTokenLocal(accessToken);
	if (!claims) return null;

	setAccessTokenCookie(event.cookies, accessToken);

	return { id: claims.sub as string };
}

export const handle: Handle = async ({ event, resolve }) => {
	const routeAccess = getRouteAccess(event.route.id ?? undefined);

	if (routeAccess.requiresAuth) {
		const accessToken = event.cookies.get('access_token');
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
