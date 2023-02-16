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

        tabs = new Array(name, "illojuan", "prueba");
    });

    //closeTab will find and delete the tab who triggered the event
    // it will manipulate find and delete the element form the tabs array
    const closeTab = (e) => {
        let index = e.detail.id;

        if (index < 0) return;

        tabs.splice(index, 1);
        tabs = [...tabs];
    };
</script>

<div class="h-full flex flex-col">
    <Topbar {tabs} on:close={closeTab} />
    <Chat />
    <Navbar />
</div>
