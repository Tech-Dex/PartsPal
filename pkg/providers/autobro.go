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
)

type Autobro struct {
	URL        string
	SearchPath string
}

func (e *Autobro) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, ctx context.Context) {
	res, err := utils.HttpGet(e.URL + e.SearchPath + *productCode)
	utils.CheckGenericProviderError(err, out)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.CheckGenericProviderError(err, out)
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckGenericProviderError(err, out)

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

					store := reflect.TypeOf(*e).Name()
					productLink, _ := ls.Find(".title").Find("a").Attr("href")
					productName := ls.Find(".title").Find("h5").Text()

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
		Store:    reflect.TypeOf(*e).Name(),
		NotFound: true,
	}

	return
}
