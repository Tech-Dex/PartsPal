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

type Comnico structs.ProviderStruct

func (p *Comnico) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
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

	doc.Find(".articol").Each(func(i int, ls *goquery.Selection) {
		if found || ctx.Err() != nil {
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

			productName := ls.Find(".denumire").Text()
			productLink, _ := ls.Find(".denumire").Attr("onclick")
			productLink = strings.Replace(productLink, "ajaxUrl('", "", -1)
			productLink = strings.Replace(productLink, "', '#content'); ", "", -1)

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
