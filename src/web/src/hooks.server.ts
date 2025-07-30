import { redirect, type Handle } from '@sveltejs/kit';
import { isProtectedRoute } from '$lib/server/auth';

export const handle: Handle = async ({ event, resolve }) => {
	// let user = null;
	const token = event.cookies.get('auth');

	if (isProtectedRoute(event.route.id ?? undefined)) {
		if (token) {
			try {
				//  validate the JWT here or trust it
				// event.locals.user = await pb.verifyJWT(token);
			} catch {
				// event.locals.user = null;
				// event.cookies.delete('auth');
			}
		} else {
			// event.locals.user = null;
			return redirect(302, '/login');
		}
	}

	return resolve(event);
};
