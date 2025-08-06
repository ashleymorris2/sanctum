import { get } from 'svelte/store';
import { authToken } from '$lib/stores/tokenStore';

export async function fetchWithToken(input: RequestInfo, init: RequestInit = {}) {
	const token = get(authToken);
	if (!token) throw new Error('No access token in memory');

	return fetch(input, {
		...init,
		headers: {
			...(init.headers || {}),
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json'
		}
	});
}
