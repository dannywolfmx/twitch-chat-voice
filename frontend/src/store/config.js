import { writable } from "svelte/store"

import { IsLoggedIn } from "../../wailsjs/go/app/MainApp"


export const logged = writable(false)

IsLoggedIn().then(r => {
  console.log("Loggedado:", r)
  logged.update(() => r)
})
