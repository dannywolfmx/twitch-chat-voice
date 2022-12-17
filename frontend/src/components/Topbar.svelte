<script>
    import Tab from "./tabs/Tab.svelte";

    export let tabs = new Array();
</script>

<div class="flex bg-gray-900 gap-2 p-2">
    {#each tabs as tab, id}
        <Tab account={tab} {id} on:close />
    {/each}
</div>
