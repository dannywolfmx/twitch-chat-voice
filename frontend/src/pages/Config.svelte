<script>
    import { onMount } from "svelte";
    import { EventsEmit } from "../../wailsjs/runtime";
    import Loggin from "../components/config/Loggin.svelte";
    import Config from "../components/config/Config.svelte";
    import { IsLogged } from "../store/config";

    let isLogged;

    IsLogged.subscribe((r) => {
        isLogged = r;
    });

    onMount(() => {
        EventsEmit("OnIsLoggedIn");
    });
</script>

<div class="w-full h-full grow m-auto overflow-scroll">
    {#if isLogged}
        <Config />
    {:else}
        <Loggin />
    {/if}
</div>
