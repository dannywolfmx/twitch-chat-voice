<script lang="ts">
    import { afterUpdate } from "svelte";

    export let messages = [];

    let element;

    afterUpdate(() => {
        if (messages) scrollToBottom(element);
    });

    const scrollToBottom = async (node) => {
        node.scroll({ top: node.scrollHeight, behavior: "smooth" });
    };
</script>

<div class="grow sm:p-6 p-2 overflow-hidden scroll-smooth ">
    <div
        bind:this={element}
        class="flex flex-col hover:overflow-y-scroll overflow-hidden h-full rounded-2"
    >
        <div class="grow" />
        {#each messages as { user, text, color }}
            <div class="flex mb-2">
                <div
                    class="font-bold text-white p-2 rounded-l-2"
                    style:background-color={color}
                >
                    {user}
                </div>
                <p
                    class="text-gray-100 font-bold p-2 border-gray-100 grow bg-gray-900 rounded-r-2"
                    style:border-color={color}
                >
                    {text}
                </p>
            </div>
        {/each}
    </div>
</div>
