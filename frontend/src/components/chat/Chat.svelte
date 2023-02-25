<script lang="ts">
    import { afterUpdate, createEventDispatcher } from "svelte";

    export let messages = [];
    export let muttedUsers = [];

    let element;

    afterUpdate(() => {
        if (messages) scrollToBottom(element);
    });

    const scrollToBottom = async (node) => {
        node.scroll({ top: node.scrollHeight, behavior: "smooth" });
    };

    const dispatch = createEventDispatcher();
    //close dispatch a account struct
    const toggleMutte = (user) => {
        dispatch("mutte", {
            user: user,
        });
    };
    function isMutted(user) {
        console.log(muttedUsers);
        return muttedUsers.includes(user);
    }
</script>

<div class="grow sm:p-6 p-2 overflow-hidden scroll-smooth ">
    <div
        bind:this={element}
        class="flex flex-col hover:overflow-y-scroll overflow-hidden h-full rounded-2"
    >
        <div class="grow" />
        {#each messages as { user, text, color }}
            <div class="flex mb-2">
                <div
                    class="font-bold text-white p-2 rounded-l-2 flex"
                    style:background-color={color}
                >
                    <button
                        on:click={() => toggleMutte(user)}
                        class="rounded-2 hover:bg-gray-700 h-6 w-6
                        bg-gray-600
                        p-1 m-r-2 flex flex-col justify-center items-center"
                    >
                        {#if isMutted(user)}
                            <i
                                class="i-tabler-volume-off text-purple-200 h-full w-full"
                            />
                        {:else}
                            <i
                                class="i-tabler-volume text-purple-200 h-full w-full"
                            />
                        {/if}
                    </button>

                    {user}
                </div>
                <p
                    class="text-gray-100 font-bold p-2 border-gray-100 grow bg-gray-900 rounded-r-2"
                    style:border-color={color}
                >
                    {text}
                </p>
            </div>
        {/each}
    </div>
</div>
