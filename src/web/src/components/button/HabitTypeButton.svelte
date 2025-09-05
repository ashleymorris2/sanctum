<script lang="ts">
    import {Hash} from '@lucide/svelte';
    import {ChartNoAxesCombined} from '@lucide/svelte';

    const options = ['average', 'target'];
    let selected = options[0];

    function select(v: typeof options[number]) {
        selected = v;
    }
</script>

<div class="dropdown">
    <div tabindex="0"
         role="button"
         class="align-middle btn btn-sm btn-outline font-normal border-base-content/20 hover:bg-primary-content/5">
        <Hash class="w-4 h-4"/>
        <span class="text-sm">{selected === 'target' ? 'Target' : 'Average'}</span>
    </div>
    <div tabindex="0" class="dropdown-content flex-g shadow-2xl w-104">
        <div class="card bg-base-100 card-xs shadow-2xl border border-base-content/15">
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
                                <span class="text-xl opacity-40">
                                   {#if v === 'target'}
                                       <Hash/>
                                   {:else}
                                       <ChartNoAxesCombined/>
                                   {/if}
                                </span>
                                <span>
                                    <span class="font-bold text-sm">
                                        {v === 'target' ? 'Count' : 'Average'}
                                    </span>
                                    <br>
                                    <span class="opacity-70 text-sm">
                                        {v === 'target' ? 'How many times/units in a time period?' : 'Keep an average over a time period'}
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