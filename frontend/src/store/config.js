import { writable } from "svelte/store"
import { GetConfig, SaveTwitchToken, SaveAnonymousUsername } from "../../wailsjs/go/repo/repoConfigFile"
import { ConnectWithTwitch } from "../../wailsjs/go/app/ConnectWithTwitch"


export const IsLogged = writable(false)
export const Config = writable({
  clientID: "",
  lang: "",
  username: "",
  token: ""
})

refreshConfig()

export const Logout = () => {
  IsLogged.update(() => false)

  Promise.all([
    SaveAnonymousUsername(""),
    SaveTwitchToken(""),
  ]).then(() => {
    refreshConfig()
  })
}

export const ConnectAnonymous = (username) => {
  SaveAnonymousUsername(username).then(() => {
    console.log("Prueba")
    refreshConfig()
  })
}

export const ConnectTwitch = () => {
  ConnectWithTwitch().then(connected => {
    if (!connected) {
      //log error
      return
    }
    refreshConfig()
  })
}

function refreshConfig() {
  GetConfig().then(r => {
    //check if the user is logged
    IsLogged.update(() => isLoggedIn(r))
    Config.update(() => r)
  })
}

function isLoggedIn({ username, token }) {
  console.log(username)
  return username != "" || token != ""
}
