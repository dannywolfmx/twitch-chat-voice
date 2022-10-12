<script>
    import { EventsEmit } from "../../wailsjs/runtime";

    //Real data
    let channel = "";
    let configDefault = true;

    //Test data
    //let channel = "dannywolfmx2";
    //let configDefault = false;

    function save() {
        EventsEmit("OnSaveConfig", channel);
        configDefault = false;
    }
</script>

<div class="grow flex flex-col gap-8">
    {#if configDefault}
        <input
            type="text"
            id="first_name"
            class=" bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            placeholder="Twitch channel"
            bind:value={channel}
            required
        />
    {:else}
        <h1 class="color-white font-bold">{channel}</h1>
    {/if}
</div>

<button
    class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
    on:click={save}
>
    Guardar
</button>
