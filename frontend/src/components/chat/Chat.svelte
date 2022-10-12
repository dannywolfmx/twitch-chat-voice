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

<div
    bind:this={element}
    class=" bg-gray-200 grow overflow-y-scroll scroll-smooth"
>
    {#each messages as { user, text, color }}
        <div
            class="flex flex-col pa-1 gap-1 border-b-4 border-gray-100 pl-8"
            style:border-color={color}
        >
            <p class="font-bold" style:color>{user}</p>
            <p class="text-black font-bold">{text}</p>
        </div>
    {/each}
</div>
