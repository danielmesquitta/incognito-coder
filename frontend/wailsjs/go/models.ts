export namespace entity {
	
	export class Solution {
	    thoughts?: string;
	    code?: string;
	    time_complexity?: string;
	    space_complexity?: string;
	
	    static createFrom(source: any = {}) {
	        return new Solution(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.thoughts = source["thoughts"];
	        this.code = source["code"];
	        this.time_complexity = source["time_complexity"];
	        this.space_complexity = source["space_complexity"];
	    }
	}

}

