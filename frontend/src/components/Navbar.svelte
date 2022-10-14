<script>
    export let selectedPage;
</script>

<div class="w-28 h-full bg-gray-600 p-4 flex flex-col gap-2">
    <button
        on:click={() => (selectedPage = "Home")}
        class="{selectedPage == 'Home'
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
    <button
        on:click={() => (selectedPage = "Home")}
        class="{selectedPage == 'Home'
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
    <button
        on:click={() => (selectedPage = "Main")}
        class="{selectedPage === 'Main'
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
    <div class="flex-grow" />
    <button
        on:click={() => (selectedPage = 4)}
        class="{selectedPage === 4
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
    <button
        on:click={() => (selectedPage = 5)}
        class="{selectedPage === 5
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
</div>
