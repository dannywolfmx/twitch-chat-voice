<script>
    import { EventsOn } from "../../../wailsjs/runtime";
    import { afterUpdate } from "svelte";

    let messages = [
        {
            text: "Prueba",
            user: "User prueba",
            color: "black",
        },
        {
            text: "Prueba 2",
            user: "User prueba",
            color: "black",
        },
    ];

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

<div class="grow p-8 overflow-hidden scroll-smooth ">
    <div
        bind:this={element}
        class="flex flex-col overflow-y-scroll h-full p-4 bg-gray-700 rounded-2 max-w-128 mx-auto"
    >
        <div class="grow" />
        {#each messages as { user, text, color }}
            <div class="flex  rounded-1 mb-2">
                <div
                    class="font-bold text-white p-2 rounded-2"
                    style:background-color={color}
                >
                    {user}
                </div>
                <p
                    class="text-gray-100 font-bold p-2  border-gray-100 grow"
                    style:border-color={color}
                >
                    {text}
                </p>
            </div>
        {/each}
    </div>
</div>
