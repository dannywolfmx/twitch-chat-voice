export namespace model {
	
	export class AnonymousUser {
	    username: string;
	
	    static createFrom(source: any = {}) {
	        return new AnonymousUser(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	    }
	}
	export class Chat {
	    name_channel: string;
	
	    static createFrom(source: any = {}) {
	        return new Chat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name_channel = source["name_channel"];
	    }
	}
	export class TwitchUser {
	    id: string;
	    broadcaster_type: string;
	    created_at: string;
	    description: string;
	    display_name: string;
	    email: string;
	    login: string;
	    profile_image_url: string;
	    offline_image_url: string;
	    view_count: number;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new TwitchUser(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.broadcaster_type = source["broadcaster_type"];
	        this.created_at = source["created_at"];
	        this.description = source["description"];
	        this.display_name = source["display_name"];
	        this.email = source["email"];
	        this.login = source["login"];
	        this.profile_image_url = source["profile_image_url"];
	        this.offline_image_url = source["offline_image_url"];
	        this.view_count = source["view_count"];
	        this.type = source["type"];
	    }
	}
	export class TwitchInfo {
	    token: string;
	    twitch_user: TwitchUser;
	
	    static createFrom(source: any = {}) {
	        return new TwitchInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.token = source["token"];
	        this.twitch_user = this.convertValues(source["twitch_user"], TwitchUser);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Config {
	    client_id: string;
	    lang: string;
	    twitch_info: TwitchInfo;
	    anonymous_user: AnonymousUser;
	    chats: Chat[];
	    mutted_users: string[];
	    samplerate_tts: number;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.lang = source["lang"];
	        this.twitch_info = this.convertValues(source["twitch_info"], TwitchInfo);
	        this.anonymous_user = this.convertValues(source["anonymous_user"], AnonymousUser);
	        this.chats = this.convertValues(source["chats"], Chat);
	        this.mutted_users = source["mutted_users"];
	        this.samplerate_tts = source["samplerate_tts"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

