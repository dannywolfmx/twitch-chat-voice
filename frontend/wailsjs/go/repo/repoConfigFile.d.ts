// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {repo} from '../models';

export function GetAnonymousUsername():Promise<string>;

export function GetClientID():Promise<string>;

export function GetConfig():Promise<repo.Config>;

export function GetLang():Promise<string>;

export function GetTwitchToken():Promise<string>;

export function GetTwitchUserInfo():Promise<repo.TwitchUser>;

export function SaveAnonymousUsername(arg1:string):Promise<void>;

export function SaveLang(arg1:string):Promise<void>;

export function SaveTwitchInfo(arg1:repo.TwitchInfo):Promise<void>;
