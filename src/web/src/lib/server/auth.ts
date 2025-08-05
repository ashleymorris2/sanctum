import { AUTH_API_BASE } from '$env/static/private';

export function isProtectedRoute(routeId?: string): boolean {
	if (!routeId) return false;

	return /^\/(auth|admin)\//.test(routeId);
}

export async function verifyAuthToken(token: string) {
	try {
		const res = await fetch(`${AUTH_API_BASE}/api/verify`, {
			method: 'POST',
			headers: { Authorization: `Bearer ${token}` }
		});
		if (!res.ok) return null;
		return await res.json(); // { userId, email, etc. }
	} catch {
		return null;
	}
}

export async function refreshAuthToken(refreshToken: string) {}
