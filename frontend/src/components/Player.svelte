<script>
    import { EventsEmit } from "../../wailsjs/runtime";
    function play() {
        EventsEmit("OnPlay", "play");
    }
</script>

<div class="flex justify-center pa-2 rounded-r-2 ">
    <button class="h-16 w-16 rounded-full mr-4 bg-gray-100" on:click={play} />
    <button class="h-16 w-16 rounded-full mr-4 bg-red-200" />
    <button class="h-16 w-16 rounded-full mr-4 bg-gray-100" />
</div>
