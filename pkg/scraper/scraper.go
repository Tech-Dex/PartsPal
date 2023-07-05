package scraper

import (
	"context"
	"github.com/Tech-Dex/PartsPal/pkg/interfaces"
	"github.com/Tech-Dex/PartsPal/pkg/providers"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"sync"
)

func FindBestDeal(bd *structs.BestDeal, productCode *string, pipe *chan *structs.Deal, wg *sync.WaitGroup, ctx context.Context) {
	for _, url := range providers.URLs {
		provider, err := providers.GetProvider(url)
		if utils.IsProviderNotFound(err) {
			continue
		}
		wg.Add(1)
		go genericCtxSearch(&provider, bd, productCode, *pipe, wg, ctx)
	}
}

func genericCtxSearch(p *interfaces.Provider, bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			(*p).Search(bd, productCode, out, ctx)
			return
		}
	}
}
