<script>
    import { EventsEmit } from "../../wailsjs/runtime";

    let pauseState = false;

    function resumeEvent() {
        EventsEmit("OnResume");
        pauseState = false;
    }

    function pauseEvent() {
        EventsEmit("OnPause");
        pauseState = true;
    }

    function nextEvent() {
        EventsEmit("OnNext");
        pauseState = false;
    }
</script>

<div class="flex justify-between rounded-r-2 grow min-w-40">
    {#if pauseState}
        <button
            class="grow rounded-l-2 bg-gray-700 p-3"
            on:click={() => resumeEvent()}
        >
            <i class="i-tabler-play text-purple-200 h-full w-full" />
        </button>
    {:else}
        <button
            class="grow rounded-l-2 bg-gray-700 p-3"
            on:click={() => pauseEvent()}
        >
            <i class="i-tabler-pause text-purple-200 h-full w-full" />
        </button>
    {/if}
    <button
        class="rounded-r-2 bg-gray-600 p-3 grow"
        on:click={() => nextEvent()}
    >
        <i class="i-tabler-player-track-next text-purple-200 h-full w-full" />
    </button>
</div>
