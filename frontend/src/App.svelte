<script>
    import { Greet } from "../wailsjs/go/main/App.js";
</script>

<main class="h-screen">
    <div class="flex h-full">
        <div class="w-1/2 bg-gray-600 rounded-8 ma-8 flex flex-col pa-8">
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
                    <!-- Dropdown menu -->
                    <div
                        id="dropdown"
                        class="hidden z-10 w-44 bg-white rounded divide-y divide-gray-100 shadow dark:bg-gray-700"
                    >
                        <ul
                            class="py-1 text-sm text-gray-700 dark:text-gray-200"
                            aria-labelledby="dropdownDefault"
                        >
                            <li>
                                <a
                                    href="#"
                                    class="block py-2 px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                                    >Dashboard</a
                                >
                            </li>
                            <li>
                                <a
                                    href="#"
                                    class="block py-2 px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                                    >Settings</a
                                >
                            </li>
                            <li>
                                <a
                                    href="#"
                                    class="block py-2 px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                                    >Earnings</a
                                >
                            </li>
                            <li>
                                <a
                                    href="#"
                                    class="block py-2 px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                                    >Sign out</a
                                >
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
            <button
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            >
                Guardar
            </button>
        </div>
        <div class="w-1/2 flex flex-col">
            <div class=" bg-gray-200 grow ma-8 overflow-y-scroll scroll-smooth">
                <div
                    class="flex flex-col pa-8 gap-1 border-b-4 border-gray-500"
                >
                    <span class="text-blue-500">mauriballes</span>
                    <span class="text-gray-900"
                        >O quiza se acelero cuando ya tenias el perro</span
                    >
                </div>
                <div
                    class="flex flex-col pa-8 gap-1 border-b-4 border-gray-500"
                >
                    <span class="text-blue-500">mauriballes</span>
                    <span class="text-gray-900"
                        >O quiza se acelero cuando ya tenias el perro</span
                    >
                </div>
                <div
                    class="flex flex-col pa-8 gap-1 border-b-4 border-gray-500"
                >
                    <span class="text-blue-500">mauriballes</span>
                    <span class="text-gray-900"
                        >O quiza se acelero cuando ya tenias el perro</span
                    >
                </div>
                <div
                    class="flex flex-col pa-8 gap-1 border-b-4 border-gray-500"
                >
                    <span class="text-blue-500">mauriballes</span>
                    <span class="text-gray-900"
                        >O quiza se acelero cuando ya tenias el perro</span
                    >
                </div>
                <div
                    class="flex flex-col pa-8 gap-1 border-b-4 border-gray-500"
                >
                    <span class="text-blue-500">mauriballes</span>
                    <span class="text-gray-900"
                        >O quiza se acelero cuando ya tenias el perro</span
                    >
                </div>
            </div>
            <div class="rounded-8 grid grid-cols-3 gap-3 pa-8 ma-8 bg-gray-600">
                <div class="h-16 rounded-full mr-4 bg-gray-100" />
                <div class="h-16 rounded-full mr-4 bg-gray-100" />
                <div class="h-16 rounded-full mr-4 bg-gray-100" />
            </div>
        </div>
    </div>
</main>
