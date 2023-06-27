package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
	"time"
)

type Autobro struct {
	URL string
}

func (e *Autobro) Scrape(bd *types.BestDeal, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	rf64 := 2.3125 * 100
	time.Sleep(2 * time.Second)
	price := bd.GetPrice()
	if price < rf64 {
		bd.Update(rf64, "Autobro", e.URL)
		out <- "Autobro"
		return
	}
	out <- ""
}
