package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
	"time"
)

type Autoeco struct {
	URL string
}

func (e *Autoeco) Scrape(bd *types.BestDeal, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	rf64 := 3.51231 * 100
	price := bd.GetPrice()
	if price < rf64 {
		bd.Update(rf64, "Autoeco", e.URL)
		out <- "Autoeco"
		return
	}
	out <- "Autoeco"
}
