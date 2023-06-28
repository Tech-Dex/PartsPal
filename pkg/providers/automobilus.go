package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Automobilus struct {
	URL        string
	SearchPath string
}

func (e *Automobilus) Search(bd *structs.BestDeal, productCode *string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Automobilus"
}
