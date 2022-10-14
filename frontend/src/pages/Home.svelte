<script>
    import Chat from "../components/chat/Chat.svelte";
    import Player from "../components/Player.svelte";
</script>

<div class=" flex flex-col rounded-2 w-full">
    <div class="grow p-8 max-w-164 self-center w-full">
        <Chat />
    </div>
    <Player />
</div>
