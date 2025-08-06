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
			headers: { Cookie: `refresh_token=${refreshToken}` },
			credentials: 'include'
		});
		if (!res.ok) return null;
		const { authToken, userId } = await res.json();
		return { authToken, userId };
	} catch {
		return null;
	}
}
