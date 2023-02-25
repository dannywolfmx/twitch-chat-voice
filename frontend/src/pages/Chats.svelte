<script>
    import { EventsOn } from "../../wailsjs/runtime";
    import Chat from "../components/chat/Chat.svelte";
    import Navbar from "../components/Navbar.svelte";
    import Topbar from "../components/Topbar.svelte";
    import {
        AddChat,
        GetChats,
        GetMuttedUsers,
        RemoveChat,
        ToggleMuttedUser,
    } from "../../wailsjs/go/usecase/config";
    import { model } from "../../wailsjs/go/models";

    let tabs = new Array();
    let chats = new Map();

    let selectedTabChannel = "";
    let messages = [];
    let muttedUsers = [];

    async function refreshTabs() {
        const chatsData = await GetChats();

        if (chatsData == null) {
            return;
        }

        tabs = chatsData.map((r) => {
            r.name_channel;
            if (!chats.has(r.name_channel)) {
                chats.set(r.name_channel, []);
            }

            return r.name_channel;
        });

        console.log(selectedTabChannel);
        updateSelectedTabChannel(selectedTabChannel);
    }

    //closeTab will find and delete the tab who triggered the event
    // it will manipulate find and delete the element form the tabs array
    const closeTab = (e) => {
        let index = e.detail.id;

        if (index < 0) return;

        RemoveChat(tabs[index]).then(() => {
            if (selectedTabChannel == tabs[index]) {
                selectedTabChannel = "";
            }

            refreshTabs();
        });
    };

    const addTab = (e) => {
        let chat = new model.Chat();
        selectedTabChannel = e.detail.name;
        chat.name_channel = selectedTabChannel;

        AddChat(chat).then(refreshTabs);
    };

    EventsOn("OnNewMessage", (data) => {
        let message = {
            text: data.Message,
            user: data.User.Name,
            color: data.User.Color,
        };

        if (chats.has(data.Channel)) {
            chats.get(data.Channel).push(message);
            if (selectedTabChannel == data.Channel) {
                messages = [...chats.get(selectedTabChannel)];
            }
        }
    });

    function updateSelectedTabChannel(channelName) {
        if (tabs.length == 0) {
            messages = [
                {
                    user: "Chat",
                    text: "Usa el tab de arriba para agregar algun chat",
                    color: "black",
                },
            ];
            return;
        }

        if (channelName == "") {
            channelName = tabs[tabs.length - 1];
        }

        selectedTabChannel = channelName;

        messages = [...chats.get(selectedTabChannel)];
    }

    function toggleMutte(event) {
        ToggleMuttedUser(event.detail.user).then((usersWithMutte) => {
            muttedUsers = [...usersWithMutte];
            if (chats.has(selectedTabChannel)) {
                messages = [...chats.get(selectedTabChannel)];
            }
        });
    }

    async function getMuttedUsers() {
        muttedUsers = await GetMuttedUsers();
    }

    // first time run page function
    refreshTabs();
    getMuttedUsers();

    //TODO: When the topbar tab are clicked on the close button, the select event are triggered too
    // we need to just follow the close event
</script>

<div class="h-full flex flex-col">
    <Topbar
        {tabs}
        on:close={closeTab}
        on:add={addTab}
        on:selectedTab={(e) => updateSelectedTabChannel(e.detail.tab)}
    />
    <Chat {messages} on:mutte={toggleMutte} {muttedUsers} />
    <Navbar />
</div>
