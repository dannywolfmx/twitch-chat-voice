<script>
    import { Continue, Next, Pause } from "../../wailsjs/go/tts/TTS";

    let pauseState = false;

    function resumeEvent() {
        Continue().then(() => {
            pauseState = false;
        });
    }

    function pauseEvent() {
        Pause().then(() => {
            pauseState = true;
        });
    }

    function nextEvent() {
        Next().then(() => {
            pauseState = false;
        });
    }
</script>

<div class="flex justify-between grow min-w-36 max-w-64">
    {#if pauseState}
        <button
            class="grow rounded-l-2 bg-gray-700 p-3 animate-pulse"
            on:click={() => resumeEvent()}
        >
            <i class="i-tabler-play text-purple-200 h-full w-full " />
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
        class="rounded-r-2 bg-gray-600 p-3 grow hover:bg-gray-700 "
        on:click={() => nextEvent()}
    >
        <i class="i-tabler-player-track-next text-purple-200 h-full w-full " />
    </button>
</div>
