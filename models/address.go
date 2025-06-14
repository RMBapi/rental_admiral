package models

type ShortAddress struct {
    EN string `json:"en"`
    BN string `json:"bn"`
}

type Address struct {
    PlaceID       int64        `json:"place_id"`
    Address       string       `json:"address"`
    Lat           float64      `json:"lat"`
    Lon           float64      `json:"lon"`
    ShortAddress  ShortAddress `json:"short_address"`
}


type SuggestedAddress struct {
    Results []Address `json:"results"`
}
