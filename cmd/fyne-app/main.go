package main

import (
	"fmt"
	"github.com/Tech-Dex/PartsPal/pkg/providers"
	"github.com/Tech-Dex/PartsPal/pkg/scraper"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"sync"
	"time"
)

const Timeout = 5 * time.Second

func main() {
	bd := &structs.BestDeal{
		Product: "Random Product",
		Price:   -1,
		Store:   "",
		Link:    "",
	}

	productCode := "27025"

	var wg sync.WaitGroup
	pipe := make(chan string, providers.SizeURLs)
	defer close(pipe)

	scraper.FindBestDeal(bd, &productCode, &pipe, &wg)

	for {
		select {
		case provider := <-pipe:
			fmt.Println(provider)
			fmt.Println(bd.Get())
		case <-time.After(Timeout):
			wg.Wait()
			return
		}
	}
}
