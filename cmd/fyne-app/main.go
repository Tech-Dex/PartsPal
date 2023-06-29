package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Tech-Dex/PartsPal/pkg/providers"
	"github.com/Tech-Dex/PartsPal/pkg/scraper"
	"github.com/Tech-Dex/PartsPal/pkg/structs"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const Timeout = 5 * time.Second

func main() {
	a := app.New()

	w := a.NewWindow("PartsPal")

	productListB := binding.BindStringList(
		&[]string{},
	)
	bestDealB := binding.NewString()
	bestDealB.Set("No best deal found")

	titleW := widget.NewLabel("PartsPal")
	searchW := widget.NewEntry()
	searchW.SetPlaceHolder("Product code")
	searchBtn := widget.NewButton("Search", func() {
		productListB.Set([]string{})
		go func() {
			bd := &structs.BestDeal{
				Product: "Random Product",
				Price:   -1,
				Store:   "",
				Link:    "",
			}

			//productCode := "27025"

			var wg sync.WaitGroup
			pipe := make(chan string, providers.SizeURLs)
			defer close(pipe)

			scraper.FindBestDeal(bd, &searchW.Text, &pipe, &wg)

			for {
				select {
				case provider := <-pipe:
					bestDealB.Set(strconv.FormatFloat(bd.GetPrice(), 'f', 2, 64))
					productListB.Append(provider)
				case <-time.After(Timeout):
					wg.Wait()
					return
				}
			}
		}()
	})
	searchC := container.New(
		layout.NewAdaptiveGridLayout(2),
		searchW,
		searchBtn)

	headerC := container.New(
		layout.NewAdaptiveGridLayout(5),
		layout.NewSpacer(),
		titleW,
		layout.NewSpacer(),
		searchC,
		layout.NewSpacer(),
	)

	productListW := widget.NewListWithData(productListB,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	productListC := container.New(
		layout.NewAdaptiveGridLayout(1),
		productListW,
	)

	bestDealLbl := widget.NewLabel("Best Deal")
	bestDealW := widget.NewLabelWithData(bestDealB)
	urlHl, _ := url.Parse("https://www.google.com")
	bestDealOpenLinkBtn := widget.NewHyperlinkWithStyle("Go to shop", urlHl, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	bestDealC := container.New(
		layout.NewAdaptiveGridLayout(1),
		bestDealLbl,
		bestDealW,
		bestDealOpenLinkBtn,
	)

	mainC := container.New(
		layout.NewAdaptiveGridLayout(2),
		productListC,
		bestDealC,
	)

	w.SetContent(container.New(
		layout.NewAdaptiveGridLayout(1),
		headerC,
		layout.NewSpacer(),
		mainC,
	))

	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()

}
