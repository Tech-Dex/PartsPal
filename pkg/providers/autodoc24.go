package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Autodoc24 struct {
	URL        string
	SearchPath string
}

func (e *Autodoc24) Search(bd *structs.BestDeal, productCode *string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	rf64 := 1.51231 * 100
	price := bd.GetPrice()
	if price > rf64 || price == -1 {
		bd.Update(rf64, "Autodoc24", e.URL)
		out <- "Autodoc24"
		return
	}
	out <- "Autodoc24"
}
