package providers

import (
	"github.com/Tech-Dex/PartsPal/pkg/interfaces"
	"github.com/Tech-Dex/PartsPal/pkg/utils"
)

func GetProvider(URL string) (interfaces.Provider, error) {
	switch URL {
	case "https://www.epiesa.ro":
		return &Epiesa{
			URL:        URL,
			SearchPath: "/cautare-piesa/?find=",
		}, nil
	case "https://www.autoeco.ro":
		return &Autoeco{URL: URL,
			SearchPath: "/cauta/?find=",
		}, nil
	case "https://www.autokarma.ro":
		return &Autokarma{URL: URL,
			SearchPath: "/cautare-dupa-cod-produs?src=",
		}, nil
	case "https://www.targuldepiese.ro":
		return &Targuldepiese{URL: URL}, nil
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
	case "https://www.topautopiese.ro":
		return &Topautopiese{URL: URL}, nil
	case "https://www.piese-auto.ro":
		return &Pieseauto{URL: URL}, nil
	case "https://www.rapidauto.ro":
		return &Rapidauto{URL: URL}, nil
	case "https://www.automobilus.ro":
		return &Automobilus{
			URL:        URL,
			SearchPath: "/cautare?search=",
		}, nil
	case "https://www.trol.ro":
		return &Trol{URL: URL}, nil
	case "https://www.ssvauto.ro":
		return &Ssvauto{URL: URL}, nil
	case "https://www.comnico.ro":
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
