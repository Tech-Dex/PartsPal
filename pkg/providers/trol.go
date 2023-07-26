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

type Trol structs.ProviderStruct

func (p *Trol) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
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

	doc.Find(".product-thumb").Each(func(i int, ls *goquery.Selection) {
		if found || ctx.Err() != nil {
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

			productLink, _ := ls.Find(".caption").Find("h4").Find("a").Attr("href")
			productName := ls.Find(".caption").Find("h4").Find("a").Text()
			if !strings.Contains(productLink, "trol.ro") {
				productLink = p.URL + productLink
			}
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
