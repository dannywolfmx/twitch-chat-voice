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

<div class="flex gap-10 items-center max-w-200 mx-auto">
    <div class="flex flex-col gap-20 grow items-center w-64">
        <div class="h-32 w-32 bg-purple-700 rounded-2 p-2">
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
    <div class="h-32 w-3 bg-gray-600 rounded" />
    <div class="flex flex-col gap-20 grow items-center w-64 ">
        <div class="h-32 w-32 bg-purple-700 rounded-2 p-2">
            <i class="i-tabler-spy h-full w-full text-purple-200" />
        </div>
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
                    class="bg-purple-600 hover:bg-purple-700 text-white font-bold py-2 px-4 rounded"
                    on:click={save}
                >
                    Conectar
                </button>
            </div>
        </div>
    </div>
</div>
