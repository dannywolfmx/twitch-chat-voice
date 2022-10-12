<script>
    import { EventsOn } from "../../../wailsjs/runtime";
    import { afterUpdate } from "svelte";

    let messages = [
        {
            text: "Prueba",
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

<p class="text-white">
    Is dark mode: {window.matchMedia("(prefers-color-scheme: dark)").matches}
</p>
<div bind:this={element} class="grow overflow-y-scroll scroll-smooth">
    {#each messages as { user, text, color }}
        <div
            class="flex gap-1 border-l-4 border-gray-100 p-3 rounded-1 mb-2"
            style:border-color={color}
        >
            <p class="font-bold text-white">{user}:</p>
            <p class="text-gray-100 font-bold">{text}</p>
        </div>
    {/each}
</div>
