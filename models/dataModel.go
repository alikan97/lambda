package models

import "github.com/google/uuid"

type MessageDTO[T any] struct {
	MessageType    string    `json:"messagetype"`
	MessageId      uuid.UUID `json:"messageid"`
	MessageContent []T       `json:"messagecontent"`
}

type AssetQuote struct {
	SymbolName string `json:"symbol"`
	Price      string `json:"price"`
}

type RawRecentTrades struct {
	Id            int64  `json:"id"`
	QuoteQuantity string `json:"quoteQty"`
	Price         string `json:"price"`
	Quantity      string `json:"qty"`
	Time          int64  `json:"time"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
	IsBestMatch   bool   `json:"isBestMatch"`
}

type RecentTradesDTO struct {
	AssetName string  `json:"assetName"`
	AssetCode string  `json:"string"`
	Price     uint32  `json:"price"`
	Quantity  float64 `json:"quantity"`
	Time      uint64  `json:"time"`
}
