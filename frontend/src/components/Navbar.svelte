<script>
    import Player from "./Player.svelte";
    export let selectedPage;
</script>

<div class="bg-gray-900 p-2 flex gap-2">
    <button
        on:click={() => (selectedPage = "Main")}
        class="{selectedPage == 'Main'
            ? 'bg-purple-900'
            : 'bg-gray-800 hover:bg-purple-900'} grow min-w-12 max-w-30 rounded-2 p-3"
    >
        <i class="i-tabler-message-circle-2 h-full w-full text-purple-200" />
    </button>
    <div class="grow" />
    <Player />
    <div class="grow" />
    <button
        on:click={() => (selectedPage = "Config")}
        class="{selectedPage === 'Config'
            ? 'bg-blue-900'
            : 'bg-gray-800 hover:bg-blue-900'} grow min-w-12 max-w-30 rounded-2 p-3"
    >
        <i class="i-tabler-settings h-full w-full text-blue-200" />
    </button>
</div>
