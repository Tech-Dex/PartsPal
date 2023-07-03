package utils

import "github.com/Tech-Dex/PartsPal/pkg/structs"

const GenericProviderErrorMsg = "Error occurred while searching"
const ProductNotFoundMsg = "Product not found"
const LaCerereMsg = "La cerere"
const IndisponibilMsg = "Indisponibil"

type ProviderNotFoundError struct {
	Provider string
}

func (e *ProviderNotFoundError) Error() string {
	return e.Provider + " not found"
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

func CheckGenericProviderError(err error, out chan<- *structs.Deal) {
	if err != nil {
		out <- &structs.Deal{Error: GenericProviderErrorMsg}
	}
}
