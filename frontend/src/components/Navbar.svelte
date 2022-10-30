<script>
    import Player from "./Player.svelte";
    export let selectedPage;
</script>

<div class="bg-gray-900 sm:p-2 flex sm:gap-2">
    <button
        on:click={() => (selectedPage = "Main")}
        class="{selectedPage == 'Main'
            ? 'bg-purple-800'
            : 'bg-gray-800 hover:bg-purple-800'} grow min-w-20 max-w-30 h-full sm:rounded-2 p-3"
    >
        <i class="i-tabler-message-circle-2 h-full w-full text-purple-200" />
    </button>
    <div class="sm:grow" />
    <Player />
    <div class="sm:grow" />
    <button
        on:click={() => (selectedPage = "Config")}
        class="{selectedPage === 'Config'
            ? 'bg-purple-800'
            : 'bg-gray-800 hover:bg-purple-800'} grow min-w-20 max-w-30 h-full sm:rounded-2 p-3"
    >
        <i class="i-tabler-settings h-full w-full text-purple-200" />
    </button>
</div>
