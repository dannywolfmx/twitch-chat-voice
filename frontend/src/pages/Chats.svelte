<script>
    import { Config } from "../store/config";
    import Chat from "../components/chat/Chat.svelte";
    import Navbar from "../components/Navbar.svelte";
    import Topbar from "../components/Topbar.svelte";
    import {
        GetChats,
        RemoveChat,
        AddChat,
    } from "../../wailsjs/go/repo/repoConfigFile";
    import { repo } from "../../wailsjs/go/models";

    let tabs = new Array();

    Config.subscribe((configs) => {
        if (configs.length == 0) return;
        let config = configs[0];

        const accountInfo = config.twitch_info;
        if (accountInfo == undefined) {
            return;
        }

        const name = accountInfo.twitch_user.display_name;
    });

    const refreshTabs = () => {
        GetChats().then((chats) => {
            if (chats != null) {
                tabs = chats.map((r) => r.name_channel);
            }
        });
    };

    //closeTab will find and delete the tab who triggered the event
    // it will manipulate find and delete the element form the tabs array
    const closeTab = (e) => {
        let index = e.detail.id;

        if (index < 0) return;

        RemoveChat(tabs[index]).then(() => {
            console.log(tabs[index]);
            refreshTabs();
        });
    };

    const addTab = (e) => {
        let chat = new repo.Chat();
        chat.name_channel = e.detail.name;
        AddChat(chat).then(refreshTabs);
    };

    refreshTabs();
</script>

<div class="h-full flex flex-col">
    <Topbar {tabs} on:close={closeTab} on:add={addTab} />
    <Chat />
    <Navbar />
</div>
