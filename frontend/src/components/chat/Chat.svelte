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
        let message = {
            text: data.Message,
            user: data.User.Name,
        };
        messages = [...messages, message];
    });
</script>

<div
    bind:this={element}
    class=" bg-gray-200 grow overflow-y-scroll scroll-smooth"
>
    {#each messages as { user, text }}
        <div class="flex flex-col pa-1 gap-1 border-b-4 border-gray-100">
            <p class="text-blue-500">{user}</p>
            <p class="text-gray-900">{text}</p>
        </div>
    {/each}
</div>
