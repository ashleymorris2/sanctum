import { AUTH_API_BASE } from '$env/static/private';

export async function verifyAuthToken(token: string) {
	try {
		const res = await fetch(`${AUTH_API_BASE}/api/verify`, {
			method: 'POST',
			headers: { Authorization: `Bearer ${token}` }
		});
		if (!res.ok) return null;
		return await res.json();
	} catch {
		return null;
	}
}

export async function refreshAuthToken(refreshToken: string) {
	try {
		const res = await fetch(`${AUTH_API_BASE}/api/refresh`, {
			method: 'POST',
			headers: { Cookie: `refresh_token=${refreshToken}` }, // Note: cookies only work client-side
			credentials: 'include'
		});
		if (!res.ok) return null;

		const { token, user } = await res.json();
		return { token, user };
	} catch {
		return null;
	}
}
