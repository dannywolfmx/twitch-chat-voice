<script>
    export let selectedPage;
</script>

<div class="w-20 h-full bg-gray-900 p-2 flex flex-col gap-2">
    <button
        on:click={() => (selectedPage = "Main")}
        class="{selectedPage == 'Main'
            ? 'bg-purple-800'
            : 'bg-gray-800 hover:bg-purple-800'} w-full h-16 rounded-2 p-3"
    >
        <i
            class="i-tabler-brand-twitch h-full w-full text-purple-200"
        /></button
    >
    <div class="flex-grow" />
    <button
        on:click={() => (selectedPage = "Config")}
        class="{selectedPage === 'Config'
            ? 'bg-purple-800'
            : 'bg-gray-800 hover:bg-purple-800'} w-full h-16 rounded-2 p-3"
        ><i class="i-tabler-settings h-full w-full text-purple-200" /></button
    >
</div>
