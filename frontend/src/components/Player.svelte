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

<div class="flex justify-center pa-2 rounded-r-2 ">
    {#if pauseState}
        <button
            class="h-16 w-16 rounded-full mr-4 bg-gray-200 p-2"
            on:click={() => resumeEvent()}
        >
            <i class="i-mdi-play text-gray-800 h-full w-full" />
        </button>
    {:else}
        <button
            class="h-16 w-16 rounded-full mr-4 bg-red-200 p-2"
            on:click={() => pauseEvent()}
        >
            <i class="i-mdi-pause text-gray-800 h-full w-full" />
        </button>
    {/if}
    <button
        class="h-16 w-16 rounded-full mr-4 bg-gray-100 p-2"
        on:click={() => nextEvent()}
    >
        <i class="i-mdi-skip-next text-gray-800 h-full w-full" />
    </button>
</div>
