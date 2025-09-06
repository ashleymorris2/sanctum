<script lang="ts">
    import {Hash, ChartNoAxesCombined, Check} from '@lucide/svelte';

    const iconSizeClass = "w-5 h-5";
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
    <div tabindex="0" class="dropdown-content flex-g shadow-2xl w-96">
        <div class="card bg-base-100 card-xs shadow-2xl mt-1 border border-base-content/15">
            <div class="card-body p-2">
                <div class="flex flex-col gap-2">
                    {#each options as v, i (i)}
                        <button
                                type="button"
                                role="radio"
                                aria-checked={selected === v}
                                class={`rounded-sm p-2 text-left border transition cursor-pointer ${
                                    selected===v
                                        ? 'border-primary ring-1 ring-primary/30 bg-base-100 hover:bg-base-content/15'
                                        : 'border-transparent bg-base-100 hover:bg-base-content/15'}`
                                    }
                                on:click={() => select(v)}>
                            <span class="flex items-center justify-between w-full">
                                <span class="flex items-center gap-2">
                                    <span class="opacity-25 mr-2">
                                       {#if v === 'target'}
                                           <Hash class={iconSizeClass}/>
                                       {:else}
                                           <ChartNoAxesCombined class={iconSizeClass}/>
                                       {/if}
                                    </span>
                                    <span>
                                        <span class="font-bold text-sm">
                                            {v === 'target' ? 'Count' : 'Average'}
                                        </span>
                                        <br>
                                        <span class="opacity-60 text-sm">
                                            {v === 'target' ? 'How many times/units in a time period?' : 'Keep an average over a time period'}
                                        </span>
                                    </span>
                                </span>
                                <!-- Right check -->
                                {#if selected === v}
                                    <Check class="w-4 h-4 text-primary mr-2"/>
                                {/if}
                            </span>
                        </button>
                    {/each}
                </div>
            </div>
        </div>
    </div>
</div>