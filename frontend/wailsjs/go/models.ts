export namespace repo {
	
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
	export class TwitchUser {
	    username: string;
	    token: string;
	
	    static createFrom(source: any = {}) {
	        return new TwitchUser(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.token = source["token"];
	    }
	}
	export class Config {
	    clientID: string;
	    lang: string;
	    twitchUser: TwitchUser;
	    anonymousUser: AnonymousUser;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.clientID = source["clientID"];
	        this.lang = source["lang"];
	        this.twitchUser = this.convertValues(source["twitchUser"], TwitchUser);
	        this.anonymousUser = this.convertValues(source["anonymousUser"], AnonymousUser);
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

