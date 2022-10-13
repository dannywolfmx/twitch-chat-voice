<script>
    let selected = 1;
</script>

<div class="w-28 h-full bg-gray-600 p-4 flex flex-col gap-2">
    <button
        on:click={() => (selected = 1)}
        class="{selected == 1
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
    <button
        on:click={() => (selected = 2)}
        class="{selected == 2
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
    <button
        on:click={() => (selected = 3)}
        class="{selected === 3
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
    <div class="flex-grow" />
    <button
        on:click={() => (selected = 4)}
        class="{selected === 4
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
    <button
        on:click={() => (selected = 5)}
        class="{selected === 5
            ? 'bg-gray-800'
            : 'bg-gray-700 hover:bg-gray-900'} w-full h-16 rounded-2 "
    />
</div>
