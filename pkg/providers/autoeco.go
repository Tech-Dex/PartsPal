package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Autoeco struct {
	URL        string
	SearchPath string
}

func (e *Autoeco) Search(bd *structs.BestDeal, productCode *string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	rf64 := 3.51231 * 100
	price := bd.GetPrice()
	if price > rf64 || price == -1 {
		bd.Update(rf64, "Autoeco", e.URL)
		out <- "Autoeco"
		return
	}
	out <- "Autoeco"
}
