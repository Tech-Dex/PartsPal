package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/interfaces"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
)

var URLs = []string{
	"https://www.epiesa.ro",
	"https://www.autoeco.ro",
	"https://www.autokarma.ro",
	"https://www.targuldepiese.ro",
	"https://www.autopiesa.ro",
	"https://www.autohut.ro",
	"https://www.autobro.ro",
	"https://www.automag.ro",
	"https://www.topautopiese.ro",
	"https://www.piese-auto.ro",
	"https://www.rapidauto.ro",
	"https://www.automobilus.ro",
	"https://www.trol.ro",
	"https://comnico.ro",
}

var SizeURLs = len(URLs)

func GetProvider(URL string) (interfaces.Provider, error) {
	switch URL {
	case "https://www.epiesa.ro":
		return &Epiesa{
			URL:        URL,
			SearchPath: "/cautare-piesa/?find=",
		}, nil
	case "https://www.autoeco.ro":
		return &Autoeco{
			URL:        URL,
			SearchPath: "/cauta/?find=",
		}, nil
	case "https://www.autokarma.ro":
		return &Autokarma{
			URL:        URL,
			SearchPath: "/cautare-dupa-cod-produs?src=",
		}, nil
	case "https://www.targuldepiese.ro":
		return &Targuldepiese{
			URL:        URL,
			SearchPath: "/cautare-piesa/?find=",
		}, nil
	case "https://www.autopiesa.ro":
		return &Autopiesa{
			URL:        URL,
			SearchPath: "/cautare-piese-auto?search=",
		}, nil
	case "https://www.autohut.ro":
		return &Autohut{
			URL:        URL,
			SearchPath: "/cautare-piesa/?find=",
		}, nil
	case "https://www.autobro.ro":
		return &Autobro{
			URL:        URL,
			SearchPath: "/cautare-piese-auto?search=",
		}, nil
	case "https://www.automag.ro":
		return &Automag{
			URL:        URL,
			SearchPath: "/cautare?search=",
		}, nil
	case "https://www.piese-auto.ro":
		return &Pieseauto{
			URL:        URL,
			SearchPath: "/cautare-piesa/?find=",
		}, nil
	case "https://www.rapidauto.ro":
		return &Rapidauto{
			URL:        URL,
			SearchPath: "/ro/searchresult.html?search=",
		}, nil
	case "https://www.automobilus.ro":
		return &Automobilus{
			URL:        URL,
			SearchPath: "/cautare?search=",
		}, nil
	case "https://www.trol.ro":
		return &Trol{
			URL:        URL,
			SearchPath: "/index.php?route=product/search&search=",
		}, nil
	case "https://comnico.ro":
		return &Comnico{
			URL:        URL,
			SearchPath: "/cauta/",
		}, nil
	default:
		return nil, &utils.ProviderNotFoundError{
			Provider: URL,
		}
	}
}
