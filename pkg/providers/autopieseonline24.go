package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Autopieseonline24 struct {
	URL        string
	SearchPath string
}

func (e *Autopieseonline24) Search(bd *structs.BestDeal, productCode *string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Autopieseonline24"
}
