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

type Comnico struct {
	URL        string
	SearchPath string
}

func (e *Comnico) SearchCtx(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, wg *sync.WaitGroup, ctx *context.Context) {
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

func (e *Comnico) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal) {
	res, err := utils.HttpGet(e.URL + e.SearchPath + *productCode)
	utils.CheckGenericProviderError(err, out)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.CheckGenericProviderError(err, out)
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckGenericProviderError(err, out)

	found := false

	doc.Find(".articol").Each(func(i int, ls *goquery.Selection) {
		if found {
			return
		}

		productCodeProvider := ls.Find(".codlinie").Text()
		productCodeProvider = strings.ReplaceAll(productCodeProvider, " ", "")
		if productCodeProvider == *productCode {
			priceText := ls.Find(".pretunic").Text()
			priceText = priceText[0 : len(priceText)-5] // remove " Lei"
			priceText = priceText[:len(priceText)-2] + "." + priceText[len(priceText)-2:]
			priceText = strings.ReplaceAll(priceText, ",", ".")
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			store := reflect.TypeOf(*e).Name()

			productName := ls.Find(".denumire").Text()
			productLink, _ := ls.Find(".denumire").Attr("onclick")
			productLink = strings.Replace(productLink, "ajaxUrl('", "", -1)
			productLink = strings.Replace(productLink, "', '#content'); ", "", -1)
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
