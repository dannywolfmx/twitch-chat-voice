<script>
    import routes from "./routes";
    import { EventsOn } from "../wailsjs/runtime";

    EventsOn("saludo", (text) => {
        console.log(text);
    });

    let selectedPage = "Chats";

    let renderPage = routes[selectedPage];

    $: renderPage = routes[selectedPage];
</script>

<main class="bg-gray-800 h-screen flex-col flex">
    <svelte:component this={renderPage} />
</main>
