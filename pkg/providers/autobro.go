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

type Autobro structs.ProviderStruct

func (p *Autobro) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
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

	doc.Find(".list-items").Each(func(i int, ls *goquery.Selection) {
		if found || ctx.Err() != nil {
			return
		}
		details := ls.Find(".table").Find("tbody").Find("tr")
		details.Each(func(i int, ts *goquery.Selection) {
			detailsTh := ts.Find("th").Text()
			if detailsTh == "Cod produs" {
				detailsTd := ts.Find("td").Text()
				detailsTd = strings.ReplaceAll(detailsTd, " ", "")
				if detailsTd == *productCode {
					priceText := ls.Find(".price.hidden-xs").Find("span").Text()
					priceText = priceText[0 : len(priceText)-4] // remove " RON"
					price, _ := strconv.ParseFloat(priceText, 64)

					bdPrice := bd.GetPrice()

					productLink, _ := ls.Find(".title").Find("a").Attr("href")
					productName := ls.Find(".title").Find("h5").Text()
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
			}
		})
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
