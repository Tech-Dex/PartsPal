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

func (bd *BestDeal) Update(price float64, store string, link string) {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	bd.Price = price
	bd.Store = store
	bd.Link = link
}

func (bd *BestDeal) SetProduct(product string) {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	bd.Product = product
}

func (bd *BestDeal) GetProduct() string {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	return bd.Product
}

func (bd *BestDeal) SetPrice(price float64) {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	bd.Price = price
}

func (bd *BestDeal) GetPrice() float64 {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	return bd.Price
}

func (bd *BestDeal) SetStore(store string) {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	bd.Store = store
}

func (bd *BestDeal) GetStore() string {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	return bd.Store
}

func (bd *BestDeal) SetLink(link string) {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	bd.Link = link
}

func (bd *BestDeal) GetLink() string {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	return bd.Link
}

type Deal struct {
	Product string
	Price   float64
	Store   string
	Link    string
	Error   string
}

func (d *Deal) Set(product string, price float64, store string, link string) {
	d.Product = product
	d.Price = price
	d.Store = store
	d.Link = link
}

func (d *Deal) Get() (string, float64, string, string, string) {
	return d.Product, d.Price, d.Store, d.Link, d.Error
}
