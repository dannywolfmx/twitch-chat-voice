<script>
    import { onMount } from "svelte";
    import { EventsOn, EventsEmit } from "../../wailsjs/runtime";
    import Loggin from "../components/config/Loggin.svelte";
    import Config from "../components/config/Config.svelte";
    let logged = false;

    EventsOn("IsLoggedIn", (data) => {
        logged = data;
    });

    onMount(() => {
        EventsEmit("OnIsLoggedIn");
    });
</script>

<div class="w-full h-full grow m-auto overflow-scroll">
    {#if logged}
        <Config />
    {:else}
        <Loggin />
    {/if}
</div>
