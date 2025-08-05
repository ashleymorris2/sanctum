type AccessConfig = {
	requiresAuth: boolean;
	role?: 'user' | 'admin';
};

const PROTECTED_PATHS = new Map<string, { role?: 'user' | 'admin' }>([
	['/(admin)', { role: 'admin' }],
	['/(auth)', { role: 'user' }]
]);

export function getRouteAccess(routeId?: string): AccessConfig {
	if (!routeId) return { requiresAuth: false };

	for (const [path, config] of PROTECTED_PATHS) {
		if (routeId.startsWith(path)) {
			return { requiresAuth: true, role: config.role };
		}
	}

	return { requiresAuth: false };
}
