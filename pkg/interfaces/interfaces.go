package interfaces

import (
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
)

type Provider interface {
	Scrape(bd *types.BestDeal, out chan<- string, wg *sync.WaitGroup)
}
