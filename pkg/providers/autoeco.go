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

type Autoeco structs.ProviderStruct

func (p *Autoeco) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
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

	doc.Find(".col-sm-6").Each(func(i int, ls *goquery.Selection) {
		if found || ctx.Err() != nil {
			return
		}
		productSku := ls.Find(".sku").Text()
		productSku = strings.ReplaceAll(productSku, " ", "")
		if productSku == *productCode {
			priceText := ls.Find(".regular-price").Text()
			priceText = priceText[0 : len(priceText)-4] // remove " RON"
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			productLink, _ := ls.Find(".prod-name").Find("a").Attr("href")
			productName := ls.Find(".prod-name").Find("a").Text()

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
