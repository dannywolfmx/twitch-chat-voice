<script>
    import { EventsEmit } from "../../../wailsjs/runtime";

    //Real data
    //let channel = "";
    //let configDefault = true;

    //Test data
    let channel = "dannywolfmx2";

    function save() {
        EventsEmit("OnSaveConfig", channel);
    }
</script>

<div class="grow flex flex-col gap-20 items-center">
    <div class="h-56 w-56 bg-purple-500 rounded" />
    <button
        class="w-full bg-purple-500 hover:bg-purple-700 text-white font-bold py-2 px-4 rounded"
        on:click={save}
    >
        Conectar con twitch
    </button>
    <div class="w-full h-3 bg-gray-600 rounded" />
    <div class="flex gap-2 w-full">
        <input
            type="text"
            id="first_name"
            class=" bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded focus:ring-blue-500 focus:border-blue-500 block p-2.5 grow"
            placeholder="Twitch channel"
            bind:value={channel}
            required
        />
        <div>
            <button
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                on:click={save}
            >
                Conectar
            </button>
        </div>
    </div>
</div>
