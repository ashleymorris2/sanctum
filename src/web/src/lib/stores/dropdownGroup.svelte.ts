// A small helper to manage the state of a group of dropdowns.
// Ensures only one dropdown in the group can be open at a time.
// export function createDropdownGroup() {
// 	// The ID of the currently open dropdown.
// 	// - string = the id of the open dropdown
// 	// - null   = no dropdowns are open
// 	// const openId = writable<string | null>(null);
// 	let openId = $state<string | null>(null);

// 	// Toggle a dropdown open/closed:
// 	// - If the given id is already open → close it (set null).
// 	// - If a different id is open → switch to this one.
// 	function toggle(id: string) {
// 		openId = openId === id ? null : id;
// 	}

// 	// Close all dropdowns in the group.
// 	function close() {
// 		openId = null;
// 	}

// 	return { openId, toggle, close };
// }

export class DropdownGroup {
	// currently-open id, or null if none
	openId = $state<string | null>(null);

	// open a specific id
	open = (id: string) => {
		this.openId = id;
	};

	get value() {
		return this.openId;
	}

	set value(value: string | null) {
		this.openId = value;
	}

	// toggle an id (open/close)
	toggle = (id: string) => {
		this.openId = this.openId === id ? null : id;
	};

	// close all
	close = () => {
		this.openId = null;
	};
}
