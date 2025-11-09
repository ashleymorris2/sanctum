import { AUTH_API_BASE } from '$env/static/private';
import { createRemoteJWKSet, jwtVerify, type JWTPayload } from 'jose';

const JWKS = createRemoteJWKSet(new URL(`${AUTH_API_BASE}/.well-known/jwks.json`));

export async function verifyAccessTokenLocal(token: string): Promise<JWTPayload | null> {
	try {
		const { payload } = await jwtVerify(token, JWKS, {
			issuer: 'metrics-api',
			audience: 'metrics-web',
			clockTolerance: 5 // seconds of leeway
		});
		return payload;
	} catch {
		return null;
	}
}

export async function verifyAccessTokenRemote(token: string) {
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
		const { accessToken } = await res.json();
		return accessToken;
	} catch {
		return null;
	}
}
