// A small helper to manage the state of a group of dropdowns.
// Ensures only one dropdown in the group can be open at a time.
export function createDropdownGroup() {
	// The ID of the currently open dropdown.
	// - string = the id of the open dropdown
	// - null   = no dropdowns are open
	// const openId = writable<string | null>(null);
	let openId = $state<string | null>(null);

	// Toggle a dropdown open/closed:
	// - If the given id is already open → close it (set null).
	// - If a different id is open → switch to this one.
	function toggle(id: string) {
		openId = openId === id ? null : id;
	}

	// Close all dropdowns in the group.
	function close() {
		openId = null;
	}

	return { openId, toggle, close };
}
