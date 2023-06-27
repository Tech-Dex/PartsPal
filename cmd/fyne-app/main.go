package main

import (
	"github.com/Tech-Dex/PartsPal/pkg/scraper"
	"github.com/Tech-Dex/PartsPal/pkg/types"
)

func main() {
	bd := &types.BestDeal{
		Product: "Random Product",
		Price:   0.0,
		Store:   "",
		Link:    "",
	}
	scraper.Scrape(bd)
	//var wg sync.WaitGroup
	//for range make([]int, 10) {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		rf64 := rand.Float64() * 100
	//		price, _, _ := bd.Get()
	//		if rf64 > price {
	//			bd.Update(rf64, "store", "link")
	//		}
	//	}()
	//}
	//
	//wg.Wait()
	//price, store, link := bd.Get()
	//fmt.Printf("Best deal for %s is %f at %s (%s)\n", bd.Product, price, store, link)
}
