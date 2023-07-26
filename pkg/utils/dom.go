package utils

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"io"
)

func GenericGoQueryDoc(p *structs.ProviderStruct, productCode *string, out chan<- *structs.Deal) *goquery.Document {
	res, err := HttpGet(p.URL + p.SearchPath + *productCode)
	CheckGenericProviderError(err, p, out)

	if res == nil {
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		CheckGenericProviderError(err, p, out)
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckGenericProviderError(err, p, out)

	return doc
}
