import { writable } from "svelte/store"
import { GetConfig, SaveTwitchInfo, SaveAnonymousUsername } from "../../wailsjs/go/repo/repoConfigFile"
import { ConnectWithTwitch } from "../../wailsjs/go/app/ConnectWithTwitch"
import { repo } from "../../wailsjs/go/models"


export const IsLogged = writable(false)
export const Config = writable(new repo.Config())

refreshConfig()

export const Logout = () => {
  IsLogged.update(() => false)

  Promise.all([
    SaveAnonymousUsername(""),
    SaveTwitchInfo(new repo.TwitchInfo()),
  ]).then(() => {
    refreshConfig()
  })
}

export const ConnectAnonymous = (username) => {
  SaveAnonymousUsername(username).then(() => {
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
    IsLogged.update(() => isLoggedIn(r.twitch_info))
    Config.update(() => r)
  })
}

function isLoggedIn({ token }) {
  return token != ""
}
