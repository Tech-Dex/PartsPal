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

type Automobilus structs.ProviderStruct

func (p *Automobilus) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
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

	doc.Find(".productInline").Each(func(i int, ls *goquery.Selection) {
		if found || ctx.Err() != nil {
			return
		}
		productName := strings.TrimSpace(ls.Find(".productInline__heading").Text())

		if strings.Contains(productName, *productCode) == false {
			return
		}

		productLink, _ := ls.Find(".balance-text").Attr("href")
		priceText := ls.Find(".productPricesSimple__userPrice").Find(".price").Text()
		priceText = priceText[0 : len(priceText)-5] // remove " RON"
		priceText = strings.ReplaceAll(priceText, ",", ".")
		price, _ := strconv.ParseFloat(priceText, 64)
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
