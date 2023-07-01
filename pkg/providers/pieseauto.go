package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Pieseauto struct {
	URL        string
	SearchPath string
}

func (e *Pieseauto) Search(bd *structs.BestDeal, productCode *string, out chan<- structs.Deal, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- structs.Deal{
		Store: "Pieseauto",
	}
}
