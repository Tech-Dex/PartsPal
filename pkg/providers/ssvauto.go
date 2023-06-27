package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
)

type Ssvauto struct {
	URL string
}

func (e *Ssvauto) Scrape(bd *types.BestDeal, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Ssvauto"
}
