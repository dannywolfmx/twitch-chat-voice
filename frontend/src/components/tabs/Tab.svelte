<script>
    import { createEventDispatcher } from "svelte";

    //ID of the account related to the tab
    export let id = -1;

    //Account name of the chat tab
    export let account = "";

    const dispatch = createEventDispatcher();
    //close dispatch a account struct
    const close = () => {
        dispatch("close", {
            id: id,
            account: account,
        });
    };
</script>

<div
    class="bg-gray-800 h-12 w-50 rounded-2 flex flex justify-center items-center p-2"
>
    <p class="text-purple-200 grow text-center">{account}</p>
    <button
        on:click={close}
        class="rounded-2 hover:bg-gray-700 h-6 w-6 p-1.1 flex flex-col justify-center items-center"
    >
        <i class="i-tabler-x text-purple-200 h-full w-full" />
    </button>
</div>
