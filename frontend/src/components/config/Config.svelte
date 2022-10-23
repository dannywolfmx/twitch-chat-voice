<script>
    let channel = "prueba";
    function save() {}
</script>

<div class="grid p-10 md:grid-cols-2 gap-4">
    <div
        class="flex md:flex-col gap-20 grow items-center bg-gray-900 p-4 rounded-2"
    >
        <div class="md:h-32 md:w-32 h-16 w-16 rounded-2 p-2">
            <i class="i-tabler-language h-full w-full text-gray-200" />
        </div>
        <select
            name="2"
            id="4"
            class="w-full bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5"
        >
            <option value="español">Español</option>
        </select>
    </div>
    <div
        class="flex md:flex-col gap-20 grow items-center bg-purple-900 p-4 rounded-2"
    >
        <div class="h-32 w-32 rounded-2 p-2">
            <i class="i-tabler-brand-twitch h-full w-full text-purple-200" />
        </div>
        <button
            class="w-full bg-purple-500 hover:bg-purple-700 text-white font-bold py-2 px-4 rounded"
            on:click={save}
        >
            <i class="i-tabler-brand-twitch" />
            Conectar con twitch
        </button>
    </div>
    <div
        class=" md:col-span-2 h-full flex gap-20 items-center bg-gray-900 p-4 rounded-2"
    >
        <div class="h-16 w-16 rounded-2 p-2">
            <i class="i-tabler-tool h-full w-full text-gray-200" />
        </div>
    </div>
</div>
