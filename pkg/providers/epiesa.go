package providers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type Epiesa struct {
	URL        string
	SearchPath string
}

func (e *Epiesa) Search(bd *structs.BestDeal, productCode *string, out chan<- structs.Deal, wg *sync.WaitGroup) {
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

	doc.Find(".single-sub-product").Each(func(i int, ls *goquery.Selection) {
		details := ls.Find(".sub-product-detail").First().Find("p").Text()
		productCodeProvider := strings.Split(details, "Cod producator: ")[1]
		productCodeProvider = strings.ReplaceAll(productCodeProvider, " ", "")
		if productCodeProvider == *productCode {
			priceText := ls.Find(".bricolaje-bottom-text").Text()
			priceText = priceText[0 : len(priceText)-11] // remove " Lei cu TVA"
			priceText = strings.ReplaceAll(priceText, ",", ".")
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			store := reflect.TypeOf(*e).Name()
			productLink, _ := ls.Find(".product-auto-title").Find("a").Attr("href")
			productName := ls.Find(".product-auto-title").Find("a").Find("h4").Text()

			if price < bdPrice || bdPrice == -1 {
				bd.Set(productName, price, store, e.URL+productLink)
			}

			out <- structs.Deal{
				Product: productName,
				Price:   price,
				Store:   store,
				Link:    e.URL + productLink,
			}
			found = true
			return
		}
	})

	if found {
		return
	}

	out <- structs.Deal{
		Store: reflect.TypeOf(*e).Name(),
		Error: utils.ProductNotFoundMsg,
	}

	return
}
