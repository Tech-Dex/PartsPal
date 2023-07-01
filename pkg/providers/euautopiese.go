package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Euautopiese struct {
	URL        string
	SearchPath string
}

func (e *Euautopiese) Search(bd *structs.BestDeal, productCode *string, out chan<- structs.Deal, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- structs.Deal{
		Store: "Euautopiese",
	}
}
