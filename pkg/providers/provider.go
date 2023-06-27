package providers

import (
	"errors"
	"github.com/Tech-Dex/PartsPal/pkg/interfaces"
)

func GetProvider(URL string) (interfaces.Provider, error) {
	switch URL {
	case "https://www.epiesa.ro/":
		return &Epiesa{URL: URL}, nil
	case "https://www.autoeco.ro/":
		return &Autoeco{URL: URL}, nil
	case "https://www.autokarma.ro/":
		return &Autokarma{URL: URL}, nil
	case "https://www.targuldepiese.ro/":
		return &Targuldepiese{URL: URL}, nil
	case "https://www.autodoc24.ro/":
		return &Autodoc24{URL: URL}, nil
	case "https://www.autopiesa.ro/":
		return &Autopiesa{URL: URL}, nil
	case "https://www.autohut.ro/":
		return &Autohut{URL: URL}, nil
	case "https://www.autobro.ro/":
		return &Autobro{URL: URL}, nil
	case "https://www.automag.ro/":
		return &Automag{URL: URL}, nil
	case "https://www.topautopiese.ro/":
		return &Topautopiese{URL: URL}, nil
	case "https://www.piese-auto.ro/":
		return &Pieseauto{URL: URL}, nil
	case "https://www.autopieseonline24.ro/":
		return &Autopieseonline24{URL: URL}, nil
	case "https://www.rapidauto.ro/":
		return &Rapidauto{URL: URL}, nil
	case "https://www.euautopiese.ro/":
		return &Euautopiese{URL: URL}, nil
	case "https://www.kparts.ro/":
		return &Kparts{URL: URL}, nil
	case "https://www.automobilus.ro/":
		return &Automobilus{URL: URL}, nil
	case "https://www.trol.ro/":
		return &Trol{URL: URL}, nil
	case "https://www.ssvauto.ro/":
		return &Ssvauto{URL: URL}, nil
	case "https://www.comnico.ro/":
		return &Comnico{URL: URL}, nil
	default:
		return nil, errors.New("provider not found")
	}

}
