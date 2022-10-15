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

<div
    bind:this={element}
    class="flex flex-col h-full overflow-y-scroll scroll-smooth p-8 grow"
>
    <div class="grow" />
    {#each messages as { user, text, color }}
        <div
            class="flex border-b-4 border-gray-100 rounded-1 mb-2"
            style:border-color={color}
        >
            <div
                class="font-bold text-white p-2"
                style:background-color={color}
            >
                {user}
            </div>
            <p class="text-gray-100 font-bold p-2">
                {text}
            </p>
        </div>
    {/each}
</div>
