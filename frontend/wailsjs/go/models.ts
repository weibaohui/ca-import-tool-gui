export namespace main {
	
	export class CertificateInfo {
	    alias: string;
	    subject: string;
	    issuer: string;
	    valid_from: string;
	    valid_to: string;
	
	    static createFrom(source: any = {}) {
	        return new CertificateInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.alias = source["alias"];
	        this.subject = source["subject"];
	        this.issuer = source["issuer"];
	        this.valid_from = source["valid_from"];
	        this.valid_to = source["valid_to"];
	    }
	}
	export class ImportParams {
	    file_path: string;
	
	    static createFrom(source: any = {}) {
	        return new ImportParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.file_path = source["file_path"];
	    }
	}
	export class ImportResult {
	    success: boolean;
	    message: string;
	    log: string;
	
	    static createFrom(source: any = {}) {
	        return new ImportResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.log = source["log"];
	    }
	}

}

