package utils

const GenericProviderErrorMsg = "Error occurred while searching"
const ProductNotFoundMsg = "Product not found"

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

func CheckGenericProviderError(err error, out chan<- string) {
	if err != nil {
		out <- GenericProviderErrorMsg
	}
}
