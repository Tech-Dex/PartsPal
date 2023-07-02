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
	"github.com/Tech-Dex/PartsPal/pkg/utils"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const Timeout = 10 * time.Second

func main() {
	a := app.New()

	w := a.NewWindow("PartsPal")

	providerDealListB := binding.BindStringList(
		&[]string{},
	)
	bestDealB := binding.NewString()
	bestDealB.Set("No best deal found")
	var urlHl *url.URL
	bestDealOpenLinkBtn := widget.NewHyperlinkWithStyle("Go to shop", urlHl, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	titleW := widget.NewLabel("PartsPal")
	searchW := widget.NewEntry()
	searchW.SetPlaceHolder("Product code")
	searchBtn := widget.NewButton("Search", func() {
		if searchW.Text == "" {
			return
		}
		providerDealListB.Set([]string{})
		go func() {
			bd := &structs.BestDeal{
				Product: "",
				Price:   -1,
				Store:   "",
				Link:    "",
			}

			//productCode := "27025"

			var wg sync.WaitGroup
			pipe := make(chan structs.Deal, providers.SizeURLs)
			defer close(pipe)

			scraper.FindBestDeal(bd, &searchW.Text, &pipe, &wg)

			for {
				select {
				case deal := <-pipe:
					bdProduct, bdPrice, bdStore, bdLink := bd.Get()
					if bdPrice == -1 {
						bestDealB.Set("No best deal found")
					} else {
						bestDealB.Set(bdProduct + " - " + strconv.FormatFloat(bdPrice, 'f', 2, 64) + " RON @ " + bdStore)
						bestDealOpenLinkBtn.SetURLFromString(bdLink)
					}
					_, dPrice, dStore, _, err := deal.Get()

					providerDeal := strconv.FormatFloat(dPrice, 'f', 2, 64) + " RON @ " + dStore
					if err != "" {
						providerDeal = err + " @ " + dStore
					}
					if err == utils.LaCerereMsg {
						providerDeal = "La cerere @ " + dStore
					}
					providerDealListB.Append(providerDeal)
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

	productListW := widget.NewListWithData(providerDealListB,
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
