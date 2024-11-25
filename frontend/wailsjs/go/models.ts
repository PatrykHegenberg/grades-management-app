export namespace main {
	
	export class Bewertung {
	    vorname: string;
	    nachname: string;
	    id: number;
	    hvPunkte: number;
	    hvProzent: number;
	    hvNote: number;
	    lvPunkte: number;
	    lvProzent: number;
	    lvNote: number;
	    gesamtProzent: number;
	    gesamtNote: number;
	    gewertet: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Bewertung(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.vorname = source["vorname"];
	        this.nachname = source["nachname"];
	        this.id = source["id"];
	        this.hvPunkte = source["hvPunkte"];
	        this.hvProzent = source["hvProzent"];
	        this.hvNote = source["hvNote"];
	        this.lvPunkte = source["lvPunkte"];
	        this.lvProzent = source["lvProzent"];
	        this.lvNote = source["lvNote"];
	        this.gesamtProzent = source["gesamtProzent"];
	        this.gesamtNote = source["gesamtNote"];
	        this.gewertet = source["gewertet"];
	    }
	}
	export class MaxPunkte {
	    hvMax: number;
	    lvMax: number;
	    hvGewichtung: number;
	    lvGewichtung: number;
	
	    static createFrom(source: any = {}) {
	        return new MaxPunkte(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hvMax = source["hvMax"];
	        this.lvMax = source["lvMax"];
	        this.hvGewichtung = source["hvGewichtung"];
	        this.lvGewichtung = source["lvGewichtung"];
	    }
	}

}

