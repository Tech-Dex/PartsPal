package structs

import "sync"

type BestDeal struct {
	mu      sync.Mutex
	Product string
	Price   float64
	Store   string
	Link    string
}

func (bd *BestDeal) Set(product string, price float64, store string, link string) {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	bd.Product = product
	bd.Price = price
	bd.Store = store
	bd.Link = link
}

func (bd *BestDeal) Get() (string, float64, string, string) {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	return bd.Product, bd.Price, bd.Store, bd.Link
}

func (bd *BestDeal) GetPrice() float64 {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	return bd.Price
}

type Deal struct {
	Product     string
	Price       float64
	Store       string
	Link        string
	Error       string
	NotFound    bool
	Unavailable bool
	Requestable bool
}

func (d *Deal) Set(product string, price float64, store string, link string, error string, notFound bool, unavailable bool, requestable bool) {
	d.Product = product
	d.Price = price
	d.Store = store
	d.Link = link
	d.Error = error
	d.NotFound = notFound
	d.Unavailable = unavailable
	d.Requestable = requestable
}

func (d *Deal) Get() (string, float64, string, string, string, bool, bool, bool) {
	return d.Product, d.Price, d.Store, d.Link, d.Error, d.NotFound, d.Unavailable, d.Requestable
}

type DealJson struct {
	Type        string  `json:"type"`
	Product     string  `json:"product"`
	Price       float64 `json:"price"`
	Store       string  `json:"store"`
	Link        string  `json:"link"`
	NotFound    bool    `json:"notFound"`
	Unavailable bool    `json:"unavailable"`
	Requestable bool    `json:"requestable"`
}

type ProviderStruct struct {
	URL        string
	SearchPath string
	Store      string
}
