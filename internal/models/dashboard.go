package models

type SalesCategory struct {
    Name   string `json:"name"`
    Sales  int    `json:"sales"`
    Profit int    `json:"profit"`
}

type BestSeller struct {
    ProductName string `json:"product_name"`
    Sold        int    `json:"sold"`
    Profit      int    `json:"profit"`
}

type OrderStats struct {
    OnProgress int `json:"on_progress"`
    Shipping   int `json:"shipping"`
    Done       int `json:"done"`
}