package providers

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type Automag struct {
	URL        string
	SearchPath string
}

func (e *Automag) SearchCtx(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal, wg *sync.WaitGroup, ctx *context.Context) {
	defer wg.Done()
	for {
		select {
		case <-(*ctx).Done():
			fmt.Println("Automag ctx done")
			return
		default:
			e.Search(bd, productCode, out)
			return
		}
	}
}

func (e *Automag) Search(bd *structs.BestDeal, productCode *string, out chan<- *structs.Deal) {
	res, err := utils.HttpGet(e.URL + e.SearchPath + *productCode)
	utils.CheckGenericProviderError(err, out)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.CheckGenericProviderError(err, out)
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckGenericProviderError(err, out)

	found := false

	firstFind := doc.Find(".productInline").First()
	if firstFind.Length() > 0 {
		productName := strings.TrimSpace(firstFind.Find(".productInline__heading").Text())
		productLink, _ := firstFind.Find(".balance-text").Attr("href")
		priceText := firstFind.Find(".productPricesSimple__userPrice").Find(".price").Text()

		var price float64
		if priceText != "" {
			priceText = priceText[0 : len(priceText)-5] // remove " RON"
			priceText = strings.ReplaceAll(priceText, ",", ".")
			price, _ = strconv.ParseFloat(priceText, 64)
			if price == 0 {
				out <- &structs.Deal{
					Store:       reflect.TypeOf(*e).Name(),
					Unavailable: true,
				}

				return
			}
		} else {
			out <- &structs.Deal{
				Product:     productName,
				Store:       reflect.TypeOf(*e).Name(),
				Requestable: true,
			}
			found = true
			return
		}
		bdPrice := bd.GetPrice()

		store := reflect.TypeOf(*e).Name()
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

	if found {
		return
	}

	out <- &structs.Deal{
		Store:    reflect.TypeOf(*e).Name(),
		NotFound: true,
	}

	return
}
