<script>
    let channel = "prueba";
    function save() {}
</script>

<div class="flex flex-col gap-16">
    <div class="flex gap w-full">
        <input
            type="text"
            id="first_name"
            class="font-bold bg-gray-700 text-gray-100 text-sm rounded block p-2.5 grow"
            placeholder="Twitch channel"
            bind:value={channel}
            readonly
        />
        <div>
            <button
                class="w-16 h-full bg-red-400 hover:bg-red-500  rounded-r"
                on:click={save}
            />
        </div>
    </div>
    <select
        name="2"
        id="4"
        class="w-64 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5"
    >
        <option value="español">Español</option>
    </select>
</div>
