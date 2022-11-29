<script>
    import { EventsOn } from "../../../wailsjs/runtime";
    import { afterUpdate } from "svelte";

    let messages = [];

    let element;

    afterUpdate(() => {
        if (messages) scrollToBottom(element);
    });

    const scrollToBottom = async (node) => {
        node.scroll({ top: node.scrollHeight, behavior: "smooth" });
    };

    EventsOn("OnNewMessage", (data) => {
        console.log(data.Tags);
        let message = {
            text: data.Message,
            user: data.User.Name,
            color: data.User.Color,
        };
        messages = [...messages, message];
    });
</script>

<div class="grow sm:p-8 p-2 overflow-hidden scroll-smooth ">
    <div
        bind:this={element}
        class="flex flex-col overflow-y-scroll h-full rounded-2 sm:max-w-128 mx-auto"
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
