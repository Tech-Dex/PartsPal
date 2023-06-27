package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
)

type Kparts struct {
	URL string
}

func (e *Kparts) Scrape(bd *types.BestDeal, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Kparts"
}
