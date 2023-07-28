export enum AccountType {
	MintId = "mintId",
	Img = "img",
	Expiry = "expiry",
	CurrentBid = "currentBid",
}

export interface Account {
	id: string;
	accountType: AccountType;
	tokens: number;
	callbackTimeMS: number;
	data: AccountData;
	version: number;
}

export interface AccountData {
	mintId: string;
	img: string;
	expiry: number;
	currentBid: number;
}