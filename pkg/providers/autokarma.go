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

type Autokarma structs.ProviderStruct

func (p *Autokarma) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
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

	doc.Find(".products-row").Each(func(i int, ls *goquery.Selection) {
		if found || ctx.Err() != nil {
			return
		}

		productSku := ls.Find(".sku_prod").Text()
		productSku = strings.ReplaceAll(productSku, " ", "")
		if productSku == *productCode {
			col4th := ls.Find(".col-sm-2").Eq(3)
			priceText := col4th.Find(".media-body").Find("span").Text()
			priceText = priceText[0 : len(priceText)-4] // remove " RON"
			priceText = strings.ReplaceAll(priceText, ".", "")
			priceText = strings.ReplaceAll(priceText, ",", ".")
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			productLink := p.URL + p.SearchPath + *productCode
			productName := ls.Find(".col-sm-6").First().Text()
			productName = strings.ReplaceAll(productName, "\n", "")
			productName = strings.TrimSpace(productName)
			for strings.Contains(productName, "  ") {
				productName = strings.ReplaceAll(productName, "  ", " ")
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
