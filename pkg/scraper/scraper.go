package scraper

import (
	"fmt"
	"github.com/Tech-Dex/PartsPal/pkg/providers"
	"github.com/Tech-Dex/PartsPal/pkg/types"
	"sync"
	"time"
)

func Scrape(bd *types.BestDeal) {
	var wg sync.WaitGroup
	r := make(chan string, providers.SizeURLs)
	defer close(r)
	for _, url := range providers.URLs {
		provider, err := providers.GetProvider(url)
		if err != nil && err.Error() == "provider not found" {
			continue
		}

		wg.Add(1)
		go provider.Scrape(bd, r, &wg)
	}

	for {
		select {
		case provider := <-r:
			fmt.Println(provider)
			fmt.Println(bd.Get())
		case <-time.After(5 * time.Second):
			wg.Wait()
			return
		}
	}
}
