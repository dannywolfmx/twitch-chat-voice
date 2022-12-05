import { writable } from "svelte/store"
import { GetConfig, SaveTwitchUser, SaveAnonymousUsername } from "../../wailsjs/go/repo/repoConfigFile"
import { ConnectWithTwitch } from "../../wailsjs/go/app/ConnectWithTwitch"
import { repo } from "../../wailsjs/go/models"


export const IsLogged = writable(false)
export const Config = writable(new repo.Config())

refreshConfig()

export const Logout = () => {
  IsLogged.update(() => false)

  Promise.all([
    SaveAnonymousUsername(""),
    SaveTwitchUser("", ""),
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
    IsLogged.update(() => isLoggedIn(r.twitchUser))
    Config.update(() => r)
  })
}

function isLoggedIn({ username, token }) {
  return username != "" || token != ""
}
