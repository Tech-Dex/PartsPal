package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
	"time"
)

type Autodoc24 struct {
	URL string
}

func (e *Autodoc24) Scrape(bd *types.BestDeal, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	rf64 := 1.51231 * 100
	time.Sleep(1 * time.Second)
	price := bd.GetPrice()
	if price < rf64 {
		bd.Update(rf64, "Autodoc24", e.URL)
		out <- "Autodoc24"
		return
	}
	out <- "Autodoc24"
}
