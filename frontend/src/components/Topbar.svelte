<script>
    import Tab from "./tabs/Tab.svelte";

    const closeTab = (e) => {
        console.log();
    };
</script>

<div class="flex bg-gray-900 gap-2 p-2">
    <div
        class="bg-gray-800 h-12 w-50 rounded flex-col flex justify-center items-center "
    >
        <div class="text-purple-100">Prueba</div>
    </div>
    <Tab text="Prueba 2" on:close={closeTab} />
</div>
