package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Autokarma struct {
	URL        string
	SearchPath string
}

func (e *Autokarma) Search(bd *structs.BestDeal, productCode *string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Autokarma"
}
