export namespace Computer {
	
	export class CPUInfo {
	    cpu: number;
	    vendorId: string;
	    family: string;
	    model: string;
	    stepping: number;
	    physicalId: string;
	    coreId: string;
	    cores: number;
	    modelName: string;
	    mhz: number;
	    cacheSize: number;
	    flags: string[];
	    microcode: string;
	
	    static createFrom(source: any = {}) {
	        return new CPUInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cpu = source["cpu"];
	        this.vendorId = source["vendorId"];
	        this.family = source["family"];
	        this.model = source["model"];
	        this.stepping = source["stepping"];
	        this.physicalId = source["physicalId"];
	        this.coreId = source["coreId"];
	        this.cores = source["cores"];
	        this.modelName = source["modelName"];
	        this.mhz = source["mhz"];
	        this.cacheSize = source["cacheSize"];
	        this.flags = source["flags"];
	        this.microcode = source["microcode"];
	    }
	}

}

export namespace Global {
	
	export class PCHardwork {
	    CPU: string;
	    GPU: string;
	    MainBoard: string;
	    RAM: string;
	    Power: string;
	    SSD: string[];
	    HDD: string[];
	    Chassis: string;
	    Radiator: string;
	
	    static createFrom(source: any = {}) {
	        return new PCHardwork(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CPU = source["CPU"];
	        this.GPU = source["GPU"];
	        this.MainBoard = source["MainBoard"];
	        this.RAM = source["RAM"];
	        this.Power = source["Power"];
	        this.SSD = source["SSD"];
	        this.HDD = source["HDD"];
	        this.Chassis = source["Chassis"];
	        this.Radiator = source["Radiator"];
	    }
	}

}

