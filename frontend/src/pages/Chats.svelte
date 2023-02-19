<script>
    import { EventsOn } from "../../wailsjs/runtime";
    import { Config } from "../store/config";
    import Chat from "../components/chat/Chat.svelte";
    import Navbar from "../components/Navbar.svelte";
    import Topbar from "../components/Topbar.svelte";
    import {
        GetChats,
        RemoveChat,
        AddChat,
    } from "../../wailsjs/go/usecase/config";
    import { model } from "../../wailsjs/go/models";

    let tabs = new Array();
    let chats = new Map();
    let selectedTab = "";
    let messages = [];

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
        GetChats().then((c) => {
            if (c != null) {
                tabs = c.map((r) => {
                    r.name_channel;
                    if (!chats.has(r.name_channel)) {
                        chats.set(r.name_channel, []);
                    }

                    return r.name_channel;
                });
            }
        });
    };

    //closeTab will find and delete the tab who triggered the event
    // it will manipulate find and delete the element form the tabs array
    const closeTab = (e) => {
        let index = e.detail.id;

        if (index < 0) return;

        RemoveChat(tabs[index]).then(() => {
            refreshTabs();
        });
    };

    const addTab = (e) => {
        let chat = new model.Chat();
        selectedTab = e.detail.name;
        chat.name_channel = selectedTab;

        AddChat(chat).then(refreshTabs);

        console.log(selectedTab);
        chats.set(selectedTab, []);
        messages = [...chats.get(selectedTab)];
    };

    refreshTabs();

    EventsOn("OnNewMessage", (data) => {
        let message = {
            text: data.Message,
            user: data.User.Name,
            color: data.User.Color,
        };
        console.log(data.Channel);
        chats.get(data.Channel).push(message);

        if (selectedTab == data.Channel) {
            messages = [...chats.get(selectedTab)];
        }
    });

    const updateSelectedTab = (e) => {
        selectedTab = e.detail.tab;
    };
</script>

<div class="h-full flex flex-col">
    <Topbar
        {tabs}
        on:close={closeTab}
        on:add={addTab}
        on:selectedTab={updateSelectedTab}
    />
    <Chat {messages} />
    <Navbar />
</div>
