package utils

import (
	"github.com/Tech-Dex/PartsPal/pkg/structs"
)

const GenericProviderErrorMsg = "Error occurred while searching"
const ProductNotFoundMsg = "Product not found"
const LaCerereMsg = "La cerere"
const IndisponibilMsg = "Indisponibil"

type ProviderNotFoundError struct {
	Provider string
}

func (p *ProviderNotFoundError) Error() string {
	return p.Provider + " not found"
}

func IsProviderNotFound(err error) bool {
	if err != nil {
		if _, ok := (err).(*ProviderNotFoundError); ok {
			return true
		}
		panic(err)
	}

	return false
}

func CheckGenericProviderError(err error, p *structs.ProviderStruct, out chan<- *structs.Deal) {
	if err != nil {
		out <- &structs.Deal{
			Error: GenericProviderErrorMsg,
			Link:  p.URL,
			Store: p.Store,
		}
	}
}
