<script>
    import { EventsEmit } from "../../../wailsjs/runtime";

    //Real data
    let channel = "";
    //let configDefault = true;

    //Test data
    //let channel = "dannywolfmx2";

    function save() {
        EventsEmit("OnSaveConfig", channel);
    }
</script>

<div
    class="flex lg:flex-row flex-col gap-10 items-center max-w-200 m-auto p-10 h-full "
>
    <div
        class="flex lg:flex-col gap-8 lg:gap-20 lg:grow items-center lg:w-64 bg-purple-900 p-4 rounded-2"
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
        class="flex  lg:flex-col gap-8 lg:gap-20 lg:grow items-center lg:w-64  bg-gray-900 p-4 rounded-2"
    >
        <div class="h-32 w-32 rounded-2 p-2">
            <i class="i-tabler-spy h-full w-full text-gray-200" />
        </div>
        <div class="flex w-full lg:flex-row flex-col">
            <input
                type="text"
                id="first_name"
                class="bg-gray-600  font-bold text-gray-50 text-sm lg:rounded-l-2 lg:rounded-r-0 rounded-t-2 focus:ring-blue-500 focus:border-blue-500 block p-2.5 grow"
                placeholder="Modo incógnito"
                bind:value={channel}
                required
            />
            <button
                class="bg-gray-700 hover:bg-gray-800 text-white font-bold py-2 px-4 lg:rounded-r-2 lg:rounded-l-0 rounded-b-2 grow"
                on:click={save}
            >
                Conectar
            </button>
        </div>
    </div>
</div>
