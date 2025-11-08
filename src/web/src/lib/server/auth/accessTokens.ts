import { AUTH_API_BASE } from '$env/static/private';

export async function verifyAccessToken(token: string) {
	try {
		const res = await fetch(`${AUTH_API_BASE}/api/auth/verify`, {
			method: 'POST',
			headers: { Authorization: `Bearer ${token}` }
		});
		if (!res.ok) return null;
		return await res.json();
	} catch {
		return null;
	}
}

export async function refreshAccessToken(refreshToken: string) {
	try {
		const res = await fetch(`${AUTH_API_BASE}/api/auth/refresh`, {
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
