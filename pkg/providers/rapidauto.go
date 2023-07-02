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

type Rapidauto struct {
	URL        string
	SearchPath string
}

func (e *Rapidauto) Search(bd *structs.BestDeal, productCode *string, out chan<- structs.Deal, wg *sync.WaitGroup) {
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

	doc.Find(".listing-item").Each(func(i int, ls *goquery.Selection) {
		attributesText := ls.Find(".attributes").Text()
		attributesArr := strings.Split(attributesText, "\n")
		if len(attributesArr) < 2 {
			return
		}
		productCodeProvider := attributesArr[1]
		productCodeProvider = productCodeProvider[15:]
		productCodeProvider = strings.ReplaceAll(productCodeProvider, " ", "")
		if productCodeProvider == *productCode {
			priceText := ls.Find(".item-price").Text()
			priceText = priceText[0 : len(priceText)-4] // remove " RON"
			priceText = strings.ReplaceAll(priceText, ",", ".")
			price, _ := strconv.ParseFloat(priceText, 64)

			bdPrice := bd.GetPrice()

			store := reflect.TypeOf(*e).Name()

			productName := ls.Find(".col-sm-8").Find("h3").Text()
			productLink, _ := ls.Find(".col-sm-8").Find("h3").Find("a").Attr("href")

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
