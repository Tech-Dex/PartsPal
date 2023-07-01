package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Trol struct {
	URL        string
	SearchPath string
}

func (e *Trol) Search(bd *structs.BestDeal, productCode *string, out chan<- structs.Deal, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- structs.Deal{
		Store: "Trol",
	}
}
