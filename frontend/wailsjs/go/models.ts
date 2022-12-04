export namespace repo {
	
	export class Config {
	    clientID: string;
	    lang: string;
	    username: string;
	    token: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.clientID = source["clientID"];
	        this.lang = source["lang"];
	        this.username = source["username"];
	        this.token = source["token"];
	    }
	}

}

