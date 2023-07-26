package providers

import (
	"context"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"reflect"
	"strconv"
	"strings"
)

type Automag structs.ProviderStruct

func (p *Automag) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
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

	firstFind := doc.Find(".productInline").First()
	if firstFind.Length() > 0 {
		productName := strings.TrimSpace(firstFind.Find(".productInline__heading").Text())
		productLink, _ := firstFind.Find(".balance-text").Attr("href")
		priceText := firstFind.Find(".productPricesSimple__userPrice").Find(".price").Text()

		var price float64
		if priceText != "" {
			priceText = priceText[0 : len(priceText)-5] // remove " RON"
			priceText = strings.ReplaceAll(priceText, ",", ".")
			price, _ = strconv.ParseFloat(priceText, 64)
			if price == 0 {
				out <- &structs.Deal{
					Store:       reflect.TypeOf(*p).Name(),
					Unavailable: true,
				}

				return
			}
		} else {
			out <- &structs.Deal{
				Product:     productName,
				Store:       reflect.TypeOf(*p).Name(),
				Requestable: true,
			}
			found = true
			return
		}
		bdPrice := bd.GetPrice()

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

	if found || ctx.Err() != nil {
		return
	}

	out <- &structs.Deal{
		Store:    reflect.TypeOf(*p).Name(),
		NotFound: true,
	}

	return
}
