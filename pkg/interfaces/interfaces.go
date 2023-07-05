package interfaces

import (
	"context"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
)

type Provider interface {
	Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context)
}
