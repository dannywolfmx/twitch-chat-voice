<script>
    import Tab from "./tabs/Tab.svelte";

    import { createEventDispatcher } from "svelte";
    export let tabs = [];

    const dispatch = createEventDispatcher();
    //close dispatch a account struct
    const add = () => {
        let name = prompt("Nombre del canal", "");

        if (name == null || name == "") {
            return;
        }
        dispatch("add", {
            name,
        });
    };
</script>

<div class="flex bg-gray-900 gap-2 p-2">
    {#each tabs as tab, id}
        <Tab account={tab} {id} on:close />
    {/each}
    <div
        class="bg-gray-800 h-12 w-12 rounded-2 flex flex justify-center items-center p-2"
    >
        <button
            class="rounded-2 hover:bg-gray-700 h-6 w-6 p-1.1 flex flex-col justify-center items-center"
            on:click={add}
        >
            <i class="i-tabler-plus text-purple-200 h-full w-full" />
        </button>
    </div>
</div>
