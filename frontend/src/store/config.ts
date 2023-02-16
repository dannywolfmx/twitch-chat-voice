import { writable } from "svelte/store"
import { GetConfig, SaveTwitchInfo, SaveAnonymousUsername } from "../../wailsjs/go/repo/repoConfigFile"
import { ConnectWithTwitch } from "../../wailsjs/go/app/ConnectWithTwitch"
import { repo } from "../../wailsjs/go/models"


export const IsLogged = writable(false)


export class Repository<T>{
  private data: Array<T>

  constructor() {
    this.data = new Array<T>
  }

  add(element: T): boolean {
    this.data.push(element)
    return true
  }

  setElements(elements: Array<T>): boolean {
    this.data = elements
    return true
  }

  list(): Array<T> {
    return this.data
  }
}

type fnSuscriptions<T> = (data: Array<T>) => void;

export class Usecase<T>{
  private suscriptions: Array<fnSuscriptions<T>>;
  private repository: IRepository<T>

  constructor(repository: IRepository<T>) {
    this.repository = repository
    this.suscriptions = new Array<fnSuscriptions<T>>
  }

  list(): Array<T> {
    return this.repository.list()
  }
  add(data: T) {
    this.repository.add(data)
    this.broadcast()
  }
  subscribe(fn: fnSuscriptions<T>) {
    this.suscriptions.push(fn)

    fn(this.list())
  }

  set(elements: Array<T>) {
    this.repository.setElements(elements)
    this.broadcast()
  }

  private broadcast() {
    this.suscriptions.forEach((fn) => {
      fn(this.list())
    })
  }
}

let configRepo = new Repository<repo.Config>();

export const Config = new Usecase<repo.Config>(configRepo); configRepo


export interface IRepository<T> {
  list(): Array<T>
  add(element: T): boolean
  setElements(elements: Array<T>): boolean
}


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

    Config.set([r])
  })
}

function isLoggedIn({ token }) {
  return token != ""
}
