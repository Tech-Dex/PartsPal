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

type Autoeco struct {
	URL        string
	SearchPath string
}

func (e *Autoeco) Search(bd *structs.BestDeal, productCode *string, out chan<- structs.Deal, wg *sync.WaitGroup) {
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
	doc.Find(".col-sm-6").Each(func(i int, ls *goquery.Selection) {
		if found {
			return
		}
		productSku := ls.Find(".sku").Text()
		productSku = strings.ReplaceAll(productSku, " ", "")
		if productSku == *productCode {
			priceText := ls.Find(".regular-price").Text()
			priceText = priceText[0 : len(priceText)-4] // remove " RON"
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			store := reflect.TypeOf(*e).Name()
			productLink, _ := ls.Find(".prod-name").Find("a").Attr("href")
			productName := ls.Find(".prod-name").Find("a").Text()
			if price < bdPrice || bdPrice == -1 {
				bd.Set(productName, price, store, productLink)
			}

			out <- structs.Deal{
				Product: productName,
				Price:   price,
				Store:   store,
				Link:    productLink,
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
