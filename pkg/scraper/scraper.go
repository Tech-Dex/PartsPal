package scraper

import (
	"github.com/Tech-Dex/PartsPal/pkg/providers"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"sync"
)

func FindBestDeal(bd *structs.BestDeal, productCode *string, pipe *chan string, wg *sync.WaitGroup) {
	for _, url := range providers.URLs {
		provider, err := providers.GetProvider(url)
		if utils.IsProviderNotFound(err) {
			continue
		}
		wg.Add(1)
		go provider.Search(bd, productCode, *pipe, wg)
	}
}
