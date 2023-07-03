package providers

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type Trol struct {
	URL        string
	SearchPath string
}

func (e *Trol) SearchCtx(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, wg *sync.WaitGroup, ctx *context.Context) {
	defer wg.Done()
	for {
		select {
		case <-(*ctx).Done():
			return
		default:
			e.Search(bd, productCode, out)
			return
		}
	}
}

func (e *Trol) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal) {
	res, err := utils.HttpGet(e.URL + e.SearchPath + *productCode)
	utils.CheckGenericProviderError(err, out)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.CheckGenericProviderError(err, out)
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckGenericProviderError(err, out)

	found := false

	doc.Find(".product-thumb").Each(func(i int, ls *goquery.Selection) {
		if found {
			return
		}

		caption := ls.Find(".caption").Find("p").Text()
		caption = strings.Split(caption, "Nr. articol")[1]
		productCodeProvider := caption[12:]
		productCodeProvider = strings.ReplaceAll(productCodeProvider, " ", "")

		if productCodeProvider == *productCode {
			priceText := ls.Find(".price").Text()
			priceText = strings.Split(priceText, "Fără TVA")[0]
			priceText = strings.TrimSpace(priceText)
			priceText = priceText[0 : len(priceText)-4]
			priceText = strings.ReplaceAll(priceText, ",", ".")
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			store := reflect.TypeOf(*e).Name()
			productLink, _ := ls.Find(".caption").Find("h4").Find("a").Attr("href")
			productName := ls.Find(".caption").Find("h4").Find("a").Text()

			if price < bdPrice || bdPrice == -1 {
				bd.Set(productName, price, store, e.URL+productLink)
			}
			out <- &structs.Deal{
				Product: productName,
				Price:   price,
				Store:   store,
				Link:    e.URL + productLink,
			}
			found = true
			return
		}
	})

	if found {
		return
	}

	out <- &structs.Deal{
		Store:    reflect.TypeOf(*e).Name(),
		NotFound: true,
	}

	return
}
