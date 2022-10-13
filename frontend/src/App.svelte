<script>
    import { EventsOn } from "../wailsjs/runtime";
    import Chat from "./components/chat/Chat.svelte";
    import Player from "./components/Player.svelte";
    import Form from "./components/Form.svelte";
    import Navbar from "./components/Navbar.svelte";

    EventsOn("saludo", (text) => {
        console.log(text);
    });
</script>

<main class="h-screen bg-gray-800">
    <div class="flex flex-col h-full">
        <div class="grow flex h-full ">
            <div>
                <Navbar />
            </div>
            <div class=" flex flex-col rounded-2 w-full">
                <div class="grow p-8 max-w-256">
                    <Chat />
                </div>
                <Player />
            </div>
        </div>
    </div>
</main>
