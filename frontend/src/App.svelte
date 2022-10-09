<script>
    import { EventsOn } from "../wailsjs/runtime";
    import Chat from "./components/chat/Chat.svelte";
    import Player from "./components/Player.svelte";

    EventsOn("saludo", (text) => {
        console.log(text);
    });
</script>

<main class="h-screen pa-8">
    <div class="flex h-full gap-8">
        <div class="w-1/2 bg-gray-600 rounded-2 flex flex-col pa-8">
            <div class="grow flex flex-col gap-8">
                <input
                    type="text"
                    id="first_name"
                    class=" bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="John"
                    required
                />
                <div class="flex flex-col">
                    <button
                        id="dropdownDefault"
                        data-dropdown-toggle="dropdown"
                        class="text-black bg-white hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2.5 text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                        type="button"
                        >Lenguaje <svg
                            class="ml-2 w-4 h-4"
                            aria-hidden="true"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                            xmlns="http://www.w3.org/2000/svg"
                            ><path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="2"
                                d="M19 9l-7 7-7-7"
                            /></svg
                        ></button
                    >
                </div>
            </div>
            <button
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            >
                Guardar
            </button>
        </div>

        <div class="w-1/2 flex flex-col gap-8">
            <Chat />
            <Player />
        </div>
    </div>
</main>
