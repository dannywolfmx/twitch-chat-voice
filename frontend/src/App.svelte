<script>
    import { EventsOn } from "../wailsjs/runtime";
    import Chat from "./components/chat/Chat.svelte";
    import Player from "./components/Player.svelte";
    import Form from "./components/Form.svelte";

    EventsOn("saludo", (text) => {
        console.log(text);
    });
</script>

<main class="h-screen pa-8">
    <div class="flex h-full gap-8">
        <div
            class="lg:w-180 md:w-full bg-gray-600 rounded-2 md:flex flex-col pa-8 hidden"
        >
            <Form />
        </div>

        <div class="w-full flex flex-col gap-6">
            <Chat />
            <Player />
        </div>
    </div>
</main>
