<script>
    import Loggin from "../components/config/Loggin.svelte";
    import Config from "../components/config/Config.svelte";
    let logged = true;
</script>

<div class="w-full h-full grow m-auto m-8 overflow-scroll ">
    {#if logged}
        <Loggin />
    {:else}
        <Config />
    {/if}
</div>
