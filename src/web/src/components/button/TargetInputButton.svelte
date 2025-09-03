<script lang="ts">
    import {CircleCheckBig} from '@lucide/svelte';

    let {
        target = (0),
        unit = $bindable(),
        frequency = $bindable()
    } = $props();
</script>

<div class="dropdown">
    <div tabindex="0"
         role="button"
         class="align-middle btn btn-sm btn-outline font-normal border-base-content/20 hover:bg-primary-content/5">
        <CircleCheckBig class="w-4 h-4"/>
        <span class="text-sm">{target} {unit} per {frequency === 'daily' ? 'day' : 'week'}</span>
    </div>
    <div tabindex="0" class="dropdown-content shadow-2xl">
        <div class="card bg-base-300 card-xs shadow-2xl">
            <div class="card-body p-4 ">
                <div class="flex gap-4">
                    <div class="flex flex-col w-24">
                        <label class="label">
                            <span class="text-sm">Target</span>
                        </label>
                        <input type="number"
                               bind:value={target}
                               min="1"
                               placeholder="15"
                               class="input-simple"
                        />
                    </div>
                    <div class="flex flex-col flex-1">
                        <label class="label">
                            <span class="text-sm">Unit</span>
                        </label>
                        <input type="text"
                               bind:value={unit}
                               placeholder="times"
                               class="input-simple"
                        />
                    </div>
                    <div class="flex flex-col flex-1">
                        <label class="label">
                            <span class="text-sm">Frequency</span>
                        </label>
                        <input type="text"
                               bind:value={frequency}
                               placeholder="daily"
                               class="input-simple w-full"
                        />
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>