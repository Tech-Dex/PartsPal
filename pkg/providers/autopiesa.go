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

type Autopiesa struct {
	URL        string
	SearchPath string
}

func (e *Autopiesa) SearchCtx(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, wg *sync.WaitGroup, ctx *context.Context) {
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

func (e *Autopiesa) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal) {
	res, err := utils.HttpGet(e.URL + e.SearchPath + *productCode)
	utils.CheckGenericProviderError(err, out)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.CheckGenericProviderError(err, out)
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckGenericProviderError(err, out)

	found := false

	doc.Find("#carProducts").Each(func(i int, ls *goquery.Selection) {
		if found {
			return
		}
		//find div text with classes col-xs-7 col-sm-7
		productCodeProvider := ls.Find(".col-xs-7.col-sm-7").Find("p").Eq(1).Text()[5:]
		productCodeProvider = strings.ReplaceAll(productCodeProvider, " ", "")
		if productCodeProvider == *productCode {
			priceText := ls.Find(".item_price").Text()
			priceText = priceText[0 : len(priceText)-4] // remove " Lei"
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			store := reflect.TypeOf(*e).Name()
			productLink, _ := ls.Find(".women").Find("a").Attr("href")
			productName := ls.Find(".women").Find("a").Text()

			if price < bdPrice || bdPrice == -1 {
				bd.Set(productName, price, store, productLink)
			}

			out <- &structs.Deal{
				Product: productName,
				Price:   price,
				Store:   store,
				Link:    productLink,
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
