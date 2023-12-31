package providers

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"reflect"
	"strconv"
	"strings"
)

type Autopiesa structs.ProviderStruct

func (p *Autopiesa) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
	store := reflect.TypeOf(*p).Name()

	doc := utils.GenericGoQueryDoc(&structs.ProviderStruct{
		URL:        p.URL,
		SearchPath: p.SearchPath,
		Store:      store,
	}, productCode, out)

	if doc == nil {
		return
	}

	found := false

	doc.Find("#carProducts").Each(func(i int, ls *goquery.Selection) {
		if found || ctx.Err() != nil {
			return
		}

		productCodeProvider := ls.Find(".col-xs-7.col-sm-7").Find("p").Eq(1).Text()[5:]
		productCodeProvider = strings.ReplaceAll(productCodeProvider, " ", "")
		if productCodeProvider == *productCode {
			priceText := ls.Find(".item_price").Text()

			if priceText == "" {
				out <- &structs.Deal{
					Store:       store,
					Link:        p.URL,
					Requestable: true,
				}
				found = true
				return
			}

			priceText = priceText[0 : len(priceText)-4] // remove " Lei"
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			productLink, _ := ls.Find(".women").Find("a").Attr("href")
			productName := ls.Find(".women").Find("a").Text()
			productName = strings.ReplaceAll(productName, "\n", "")

			if ctx.Err() != nil {
				found = true
				return
			}

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

	if found || ctx.Err() != nil {
		return
	}

	out <- &structs.Deal{
		Store:    reflect.TypeOf(*p).Name(),
		Link:     p.URL,
		NotFound: true,
	}

	return
}
