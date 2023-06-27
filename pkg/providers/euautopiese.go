package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
)

type Euautopiese struct {
	URL string
}

func (e *Euautopiese) Scrape(bd *types.BestDeal, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Euautopiese"
}
