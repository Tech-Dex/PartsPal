package interfaces

import (
	"context"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
)

type Provider interface {
	SearchCtx(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, wg *sync.WaitGroup, ctx *context.Context)
	Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal)
}
