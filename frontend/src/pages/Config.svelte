<script>
    import Loggin from "../components/config/Loggin.svelte";
    import Config from "../components/config/Config.svelte";
    let logged = true;
</script>

<div class="w-full m-auto m-8">
    {#if logged}
        <Loggin />
    {:else}
        <Config />
    {/if}
</div>
