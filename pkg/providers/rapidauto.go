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

type Rapidauto structs.ProviderStruct

func (p *Rapidauto) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
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

	doc.Find(".listing-item").Each(func(i int, ls *goquery.Selection) {
		if found || ctx.Err() != nil {
			return
		}

		attributesText := ls.Find(".attributes").Text()
		attributesArr := strings.Split(attributesText, "\n")
		if len(attributesArr) < 2 {
			return
		}
		productCodeProvider := attributesArr[1]
		productCodeProvider = productCodeProvider[15:]
		productCodeProvider = strings.ReplaceAll(productCodeProvider, " ", "")
		if productCodeProvider == *productCode {
			priceText := ls.Find(".item-price").Text()
			priceText = priceText[0 : len(priceText)-4] // remove " RON"
			priceText = strings.ReplaceAll(priceText, ",", ".")
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			productName := ls.Find(".col-sm-8").Find("h3").Text()
			productLink, _ := ls.Find(".col-sm-8").Find("h3").Find("a").Attr("href")

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
		NotFound: true,
	}

	return
}
