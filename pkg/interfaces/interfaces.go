package interfaces

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Provider interface {
	Search(bd *structs.BestDeal, productCode *string, out chan<- string, wg *sync.WaitGroup)
}
