<script>
    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher();
    export let text = "";

    const close = () => {
        dispatch("close");
    };
</script>

<div
    class="bg-gray-800 h-12 w-50 rounded flex flex justify-center items-center p-2"
>
    <div class="text-purple-200 grow">{text}</div>
    <button on:click={close}> X </button>
</div>
