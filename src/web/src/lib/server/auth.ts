export function isProtectedRoute(routeId?: string): boolean {
	if (!routeId) return false;

	return routeId.startsWith('/(auth)/');
	// can expand with:
	// return /(auth|admin)/.test(routeId);
}
