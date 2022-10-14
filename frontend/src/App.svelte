<script>
    import Navbar from "./components/Navbar.svelte";

    import routes from "./routes";
    import { EventsOn } from "../wailsjs/runtime";

    EventsOn("saludo", (text) => {
        console.log(text);
    });

    let selectedPage = "Home";
    let renderPage = routes[selectedPage];
    $: renderPage = routes[selectedPage];
</script>

<main class="h-screen bg-gray-800">
    <div class="flex flex-col h-full">
        <div class="grow flex h-full ">
            <div>
                <Navbar bind:selectedPage />
            </div>
            {selectedPage}
            <svelte:component this={renderPage} />
        </div>
    </div>
</main>
