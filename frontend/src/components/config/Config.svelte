<script>
    import { EventsEmit } from "../../../wailsjs/runtime";

    function connectWithTwitch() {
        EventsEmit("ConnectWithTwitch");
    }
</script>

<div class="grid sm:p-8 p-2 md:grid-cols-2 gap-4">
    <div class="flex grow items-center bg-gray-900 rounded-2">
        <div class="h-16 w-16 rounded-2 p-4">
            <i class="i-tabler-language h-full w-full text-gray-200" />
        </div>
        <select
            name="2"
            id="4"
            class="grow bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded focus:ring-blue-500 focus:border-blue-500 block py-2 px-4 h-full"
        >
            <option value="español">Español</option>
        </select>
    </div>
    <div class="flex grow items-center bg-purple-900 rounded-2">
        <div class="h-16 w-16 rounded-2 p-4">
            <i class="i-tabler-brand-twitch h-full w-full text-purple-200 " />
        </div>
        <button
            class="grow bg-purple-500 hover:bg-purple-700 text-white font-bold py-2 px-4 rounded-r h-full"
            on:click={connectWithTwitch}
        >
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
