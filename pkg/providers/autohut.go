package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
)

type Autohut struct {
	URL string
}

func (e *Autohut) Scrape(bd *types.BestDeal, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Autohut"
}
