package types

type AccountType string
const (
	Mint AccountType = "mint"
	Metadata AccountType = "metadata"
	MasterEdition AccountType = "masterEdition"
	Auction AccountType = "auction"
	AuctionData AccountType = "auctionData"
	AccountEnum AccountType = "account"
	Escrow AccountType = "escrow"
)

type Account struct {
	Id string `json:"id"`
	AccountType AccountType `json:"accountType"`
	Tokens int `json:"tokens"`
	CallbackTimeMS int `json:"callbackTimeMS"`
	Version int `json:"version"`
	Data AccountData `json:"data"`
}

type AccountData struct {
	MintID string `json:"mintId,omitempty"`
	Img string `json:"img,omitempty"`
	Expiry int `json:"expiry,omitempty"`
	CurrentBid int `json:"currentBid,omitempty"`
}