<script>
    import { EventsEmit } from "../../../wailsjs/runtime";

    import { Config, Logout } from "../../store/config";

    let display_name = "";
    let isLoggedToTwitch = false;

    Config.subscribe((c) => {
        display_name = c.twitch_info.twitch_user.display_name;
        isLoggedToTwitch = c.client_id != "";
    });

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
    {#if isLoggedToTwitch}
        <div class="flex grow items-center bg-purple-800 rounded-2">
            <div
                class="grow bg-purple-900 text-white font-bold rounded flex justify-center h-full items-center"
            >
                {display_name}
            </div>
            <button
                class="h-16 w-16 rounded-2 p-4 hover:bg-purple-900"
                on:click={Logout}
            >
                <i class="i-tabler-logout h-full w-full text-purple-200 " />
            </button>
        </div>
    {:else}
        <div class="flex grow items-center bg-purple-800 rounded-2">
            <button
                class="grow bg-purple-900 hover:bg-purple-800 text-white font-bold rounded h-full"
                on:click={connectWithTwitch}
            >
                Conectar con twitch
            </button>
            <div class="h-16 w-16 rounded-2 p-4">
                <i
                    class="i-tabler-brand-twitch h-full w-full text-purple-200 "
                />
            </div>
        </div>
    {/if}
    <div
        class=" md:col-span-2 h-full flex gap-20 items-center bg-gray-900 p-4 rounded-2"
    >
        <div class="h-16 w-16 rounded-2 p-2">
            <i class="i-tabler-tool h-full w-full text-gray-200" />
        </div>
    </div>
</div>
