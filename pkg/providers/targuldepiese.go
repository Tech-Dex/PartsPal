package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Targuldepiese struct {
	URL        string
	SearchPath string
}

func (e *Targuldepiese) Search(bd *structs.BestDeal, productCode *string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Targuldepiese"
}
