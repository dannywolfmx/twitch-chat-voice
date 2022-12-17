<script>
    import { Config } from "../store/config";
    import Chat from "../components/chat/Chat.svelte";
    import Navbar from "../components/Navbar.svelte";
    import Topbar from "../components/Topbar.svelte";

    let tabs = new Array();

    Config.subscribe((config) => {
        const accountInfo = config.twitch_info;
        if (accountInfo == undefined) {
            return;
        }

        const name = accountInfo.twitch_user.display_name;

        tabs = new Array(name);
    });

    const closeTab = (e) => {
        console.log(e.detail.id);
    };
</script>

<div class="h-full flex flex-col">
    <Topbar {tabs} on:close={closeTab} />
    <Chat />
    <Navbar />
</div>
