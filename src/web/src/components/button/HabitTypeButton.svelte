<script lang="ts">
    import {CircleCheckBig} from '@lucide/svelte';

    const options = ['average', 'target'];
    let selected = options[0];

    function select(v: typeof options[number]) {
        selected = v;
    }
</script>

<div class="dropdown flex-grow">
    <div tabindex="0"
         role="button"
         class="align-middle btn btn-sm btn-outline font-normal border-base-content/20 hover:bg-primary-content/5">
        <CircleCheckBig class="w-4 h-4"/>
        <span class="text-sm">{selected === 'target' ? 'Target' : 'Average'}</span>
    </div>
    <div tabindex="0" class="dropdown-content shadow-2xl">
        <div class="card bg-base-300 card-xs shadow-2xl">
            <div class="card-body p-4 ">
                <div class="flex gap-4">
                    {#each options as v, i (i)}
                        <button
                                type="button"
                                role="radio"
                                aria-checked={selected === v}
                                class={`rounded-2xl p-3 text-left border transition ${selected===v ? 'border-primary ring-2 ring-primary/30 bg-base-100' : 'border-base-300 bg-base-100 hover:bg-base-200'}`}
                                on:click={() => select(v)}>
                            <span class="flex items-center gap-3">
                                <span class="text-2xl">
                                    {v === 'boolean' ? '☑️' : v === 'number' ? '🔢' : '📊'}
                                </span>
                                <span>
                                    <span class="font-medium">
                                        {v === 'boolean' ? 'Yes / No' : v === 'number' ? 'Count' : 'Average'}
                                    </span>
                                    <span class="text-xs opacity-70">
                                        {v === 'boolean' ? 'Did you do it?' :
                                            v === 'number' ? 'How many times/units?' :
                                                'Keep a daily average'}
                                    </span>
                                </span>
                            </span>
                        </button>
                    {/each}
                </div>
            </div>
        </div>
    </div>
</div>