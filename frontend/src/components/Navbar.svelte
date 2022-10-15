<script>
    export let selectedPage;
</script>

<div class="w-20 h-full bg-gray-600 p-2 flex flex-col gap-2">
    <button
        on:click={() => (selectedPage = "Config")}
        class="{selectedPage === 'Config'
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
    <div class="flex-grow" />
    <button
        on:click={() => (selectedPage = "Main")}
        class="{selectedPage == 'Main'
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
</div>
