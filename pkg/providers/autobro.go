package providers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"io"
	"strconv"
	"sync"
)

type Autobro struct {
	URL        string
	SearchPath string
}

func (e *Autobro) Search(bd *structs.BestDeal, productCode *string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
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
		if found {
			return
		}

		details := ls.Find(".table").Find("tbody").Find("tr")
		details.Each(func(i int, ts *goquery.Selection) {
			detailsTh := ts.Find("th").Text()
			if detailsTh == "Cod produs" {
				detailsTd := ts.Find("td").Text()
				if detailsTd == *productCode {
					priceText := ls.Find(".price.hidden-xs").Find("span").Text()
					priceText = priceText[0 : len(priceText)-4] // remove " RON"
					price, _ := strconv.ParseFloat(priceText, 64)
					bdPrice := bd.GetPrice()
					if price < bdPrice || bdPrice == -1 {
						bd.Update(price, "Autobro", e.URL)
						out <- "Autobro"
						found = true
						return
					}
				}
			}
		})
	})

	if found {
		return
	}

	out <- utils.ProductNotFoundMsg
	return
}
