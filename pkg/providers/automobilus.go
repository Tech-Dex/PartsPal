package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
)

type Automobilus struct {
	URL string
}

func (e *Automobilus) Scrape(bd *types.BestDeal, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Automobilus"
}
